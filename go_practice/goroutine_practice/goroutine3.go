package main

import "fmt"

/*
	channel 通过数据交换实现并发的数据结构

	数据交换的主要方式:
	1. 共享内存:在不同goroutine容易发生竞争问题 ->需要Mutex加锁 -> 性能负担
	2. 通道通信:采用channel连接!!

	GO语言中的并发模型CSP模型:提倡通过通信共享内存，不通过共享内存通信

	chan 引用类型 采用make建立 返回引用 Type
	var ch1 chan int
	var ch2 chan bool
	var ch3 chan []int
	使用var建立chan 只会是<nil> -> 需要用make建立
	make(chan int, buf大小)

	channel操作: 四个
	1.make: make(chan int,1)
	2.send: ch <- 1
	3.recv: x :=<- ch， <- ch
	4.close: close(cn)
	**注意:close(ch) 会panic的时机

	补充:
	make([]chan int, 10) 是啥!!




 */

func recv(c chan int)  {
	ret := <- c
	fmt.Println("接收成功",ret)

}

/*
	无缓冲通道进行通信将导致发送和接受的goroutine同步化。
	无缓冲通道称为同步通道

	无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，
	两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到
	另一个goroutine在该通道上发送一个值


 */

func main()  {
	ch := make(chan int)
	// 需要先有人下来等!!
	go recv(ch)
	// 才能给数据
	ch <- 10
	close(ch) // 发送完了一定要关闭!!
	fmt.Println("发送成功")


}


