package main

import "fmt"


/*
	逃逸分析:决定分配内存在栈中or堆中。
	1. 分配在栈中，函数运行完可以自动将内存回收
	2. 分配在堆中，函数执行结束会交给GC处理

	分析是否放在堆中!!

	四大场景:
	1.指针逃逸
	2.栈空间不足
	3.动态类型逃逸(输入参数空接口类型)
	4.闭包引用对象逃逸
 */


type a struct {
}

func NewA() *a {
	A := new(a)
	fmt.Println("指针逃逸")
	return A

}

func closure() func() {
	closer := "闭包逃逸"
	return func() {
		closer := "逃逸成功"
		fmt.Println(closer)
	}


}



func main()  {
	NewA() // new一个新对象

	slice := make([]int, 1000, 1000)
	fmt.Println("栈空间不足", slice)
	strings := "动态逃逸"
	fmt.Printf("%v", strings)
	closure()

}