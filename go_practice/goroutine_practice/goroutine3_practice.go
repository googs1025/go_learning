package main

/*
	用context.WithCancel 方法主动关闭goroutine!!
 */

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("got the stop signal")
				return
			default:
				fmt.Println("still working")
				time.Sleep(2 * time.Second)

			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("stop the goroutine")
	cancel()
	time.Sleep(5 * time.Second)
}
