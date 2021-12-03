package main

import (
	"fmt"
	"time"
)

/*
	模拟需求: 利用slice+chan实现对多个goroutine退出操作
 */

func main() {
	channels := make([]chan struct{}, 10)

	for i := 0; i < 10; i++ {
		channels[i] = make(chan struct{})
		go worker1(channels[i])
	}

	for i , ch := range channels {
		<- ch
		fmt.Println("goroutine" , i , "退出")
	}

}


func worker1(ch chan struct{}) {
	time.Sleep(5 * time.Second)
	fmt.Println("doing something")
	ch <- struct{}{}
}
