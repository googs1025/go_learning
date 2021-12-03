package main

import (
	"fmt"
	"time"
)

/*
	可用于判断chan是否存满
 */

func main() {
	output := make(chan string, 10)

	go write(output)

	for s := range output {
		fmt.Println("通道里的结果:", s)
		time.Sleep(time.Second)
	}



}

func write(ch chan string)  {
	// 可以用for无限循环
	// select 判断是否可以放入chan
	for {
		select {
		case ch <- "hello":
			fmt.Println("可以写进去")
		default:
			fmt.Println("chan已经满了")
		}
		time.Sleep(2000 * time.Microsecond)
	}

}
