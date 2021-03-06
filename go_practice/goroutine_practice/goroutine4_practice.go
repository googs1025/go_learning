package main

import (
	"context"
	"fmt"
	"time"
)
/*
	模拟主goroutine上停止多个goroutine的需求
 */


func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, "node1")
	go worker(ctx, "node2")
	go worker(ctx, "node3")

	time.Sleep(5 * time.Second)
	fmt.Println("stop the goroutine signal")
	cancel()
	time.Sleep(3 * time.Second)

}

func worker(ctx context.Context, name string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(name, "got the stop channel")
				return
			default:
				fmt.Println(name, "still working")
				time.Sleep(2 * time.Second)


			}
		}
	}()
}
