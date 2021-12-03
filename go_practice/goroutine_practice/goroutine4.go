package main

import "fmt"

/*
	可以通过cloes()关闭chan
 */



func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// 当数据发送完后，一定要关闭通道，不然goroutine会一值阻塞!!
		close(c)
	}()

	for {
		if data, ok := <- c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main goroutine done")
}
