package main

import (
	"fmt"
	"time"
)

/*
	主动通知goroutine!!
	1.channel + select !!
	2.context

	法二:
	如果有多个goroutine 就不能都用channel 除非用 channels = make([]chan int, 10)
	放进去时: channels[i] = make(chan int) -> 一个list的位置是一个channel!!

	但是有多个goroutine或是在goroutine中有goroutine 只能用context解决
 */

func main() {

	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <- stop:
				fmt.Println("got the stop signal")
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)

			}
		}
	}()


	time.Sleep(5 * time.Second)
	// 模拟主动通知goroutine!!
	fmt.Println("stop the goroutine")
	stop<- struct{}{}
	time.Sleep(5 * time.Second)

}
