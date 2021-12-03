package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	控制并发两种常见方式
	1.WaitGroup:多个goroutine(多个job)执行同一件事情 ->被动处理!!
	2.context: 可以对goroutine做主动控制 cancel timeout 给value
 */

func main()  {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("job 1 done")
		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("job 2 done")
		wg.Done()
	}()


	wg.Wait() // 在计数器还没到0前，都会在这个检查点阻塞!!
	fmt.Println("all job done")




}
