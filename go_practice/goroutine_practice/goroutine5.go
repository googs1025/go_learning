package main

import (
	"fmt"
)

/*
	从通道循环取值!!
	当通道发送有限数据时，可以通过close关闭通道告知 通道接收值的goroutine停止等待!!
	需要用if判断chan是否close

 */


func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开goroutine将0-100数据发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <-i
		}
		close(ch1)
	}()
	// 开goroutine从ch1接受值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <- ch1 // 当通道关闭后再取值 ok=false
				if !ok {
					break
				}
				ch2 <- i*i

		}
		close(ch2)
	}()
	// main goroutine从ch2中打印
	for i := range ch2 {
		fmt.Println(i)
	}

}
