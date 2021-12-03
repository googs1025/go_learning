package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. timer基本使用
	time1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <- time1.C
	fmt.Printf("t2:%v\n", t2)

	// 2.验证timer只能响应1次
	time2 := time.NewTimer(time.Second)
	for {
		<-time2.C
		fmt.Println("时间到")
	}

	// 3.timer实现延时的功能
	time.Sleep(time.Second)

	timer3 := time.NewTimer(2*time.Second)
	<-timer3.C
	fmt.Println("2秒到")
	<-time.After(2*time.Second)
	fmt.Println("2秒到")





}
