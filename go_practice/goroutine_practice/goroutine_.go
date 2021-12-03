package main

import (
	"fmt"
	"sync"
)

/*
	进程vs线程vs协程
	1.进程是程序在os中一次的执行过程，是系统调用的基本单位
	2.线程是进程的一个执行实体，是cpu调度的基本单位，比进程小，占用资源也较少。
	3.进程中可以创建与销毁多个线程；线程也可以在进程多并发执行
	4.有独立栈空间，共享堆空间(跟逃逸分析有关)，调度由用户控制 ->用户级线程

	并发vs并行
	1.同时刻在多核心上执行 ->并行
	2.同时间段在单个核心执行 ->并发


	goroutine调度
	GPM模型
	G:GO中起的goroutine协程
	P:负责管理goroutine的队列，除了放goroutine外，还放上下文环境(函数指针、堆栈地址)，消费完去全局上取!! 或是去其他P中抢!!
	M:对系统内核的虚拟，和OS系统线程数有映射，是执行goroutine的地方

	P和M一一对应，但是P可以通过runtime.GOMAXPROCS设定。
	P管理一组G挂载在M上执行，当一个G长久阻塞，在一个M上时，runtime会new一个M，阻塞G所在的P会把其他G挂载到新的M上。当旧的G阻塞完后或是已经死掉后回收旧的M


 */

func hello() {
	fmt.Println("hello world!!")
}

var wg sync.WaitGroup

func hello1() {
	defer wg.Done()
	fmt.Println("hello11 world!!")
}


func main() {
	/*
		在main goroutine起的goroutine会跟着主人死了而死，并不会等执行结束
		如果需要等待执行结束，需要用病发原语

	 */
	// 区别!!
	hello() // 属于串行调用 !!(在main goroutine中调用)
	go hello() // 另外起一个goroutine调用 !!

	// 起多个goroutine
	// goroutine是随机性的!!
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go hello1()
	}

	wg.Wait()
	fmt.Println("main goroutine done")



}
