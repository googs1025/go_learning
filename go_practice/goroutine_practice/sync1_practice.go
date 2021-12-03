package main

import (
	"fmt"
	"sync"
)

/*

	sync.WaitGroup

	wg.Add() 计数器 + delta
	wg.Done() 计数器-1
	wg.Wait() 阻塞直到计数器为0

 */


var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("Hello Goroutine")
}


func main() {

	wg.Add(1)
	go hello()
	wg.Wait()
	fmt.Println("main goroutine done!")

}
