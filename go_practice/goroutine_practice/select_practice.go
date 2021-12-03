package main

import (
	"fmt"
	"time"
)

/*
	select多路复用:
	使用场景:同时需要从多个chan接收数据时使用 ->如果没有数据接收会发生阻塞

	法一:
	for {
		尝试从ch1拿数据
		data, ok := <-ch1
		尝试从ch2拿数据
		data, ok := <-ch2
	}
	性能差， 因此使用select

	select {

	case <-chan1:
		// 如果从chan1成功读到数据，则执行此逻辑
	case chan2 <-1:
		// 如果能像chan2发数据，则执行此逻辑
	default:
		// 都没成功，执行此逻辑

	}

	场景:select同时监听一个或多个chan，直到其中一个chan ready


 */

func test1(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "test1"
}

func test2(ch chan string)  {
	time.Sleep(6 * time.Second)
	ch <- "test2"

}


func main() {

	output1 := make(chan string)
	output2 := make(chan string)

	go test1(output1)
	go test2(output2)

	select {
	case f1 := <-output1:
		fmt.Println("test01",f1)
	case f2 := <-output2:
		fmt.Println("test02", f2)
	//default:
	//	fmt.Println("还没传进来")

	}



}
