package main

import (
	"fmt"
	"runtime"
)

/*
	runtime包:
	runtime.Gosched():让出cpu时间片，重新等待安排任务

 */

func main() {
	go func(s string) {
		for i := 0; i < 10; i++ {
			fmt.Println(s)
		}

	}("world")

	for i := 0; i < 10; i++ {
		// 先把上面的走完，让出现在的时间片!
		runtime.Gosched()
		fmt.Println("hello")
	}



}
