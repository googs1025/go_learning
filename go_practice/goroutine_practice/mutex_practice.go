package main

import (
	"fmt"
	"sync"
)

/*
	会存在多个goroutine同时操作一个资源(内存地址)(临界区) ->竞争问题


 */

var x int // 全局变量
var mu sync.Mutex
var wg1 sync.WaitGroup

func add() {
	for i := 0; i < 1000; i++ {
		// 加上锁!!
		mu.Lock()
		x += 1
		mu.Unlock()

	}
	wg1.Done()
}

func main() {
	wg1.Add(2)
	go add()
	go add()

	wg1.Wait()
	fmt.Println(x)
	fmt.Println("main goroutine done")

	
}
