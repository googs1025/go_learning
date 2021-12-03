package main

import (
	"fmt"
)

/*
	多个chan同时ready，随机选择一个执行
	go func 可以写成匿名函数或是在外面的函数
 */

func main() {
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	go func() {
		int_chan <- 1

	}()

	go func() {
		string_chan <- "aaaa"
	}()

	select {
	case s1:= <-int_chan:
		fmt.Println("s1", s1)
	case s2:= <-string_chan:
		fmt.Println("s2", s2)

	}

}
