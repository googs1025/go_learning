package main

import "fmt"


/*
	闭包复制的是原对象的地址。

 */

// add函数的返回是一个函数 (闭包)
func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}

}

func test01(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并且返回

	add := func(i int) int {
		base += i
		return base

	}

	sub := func(i int) int {
		base -= i
		return base

	}

	return add, sub

}


func main()  {
	// 从外部引用函数参数局部变量
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))

	tmp2 := add(100)
	fmt.Println(tmp2(2),tmp2((5)))


	// 注意一下base的变化!!
	f1, f2 := test01(100)
	fmt.Println(f1(1),f2(2))
	fmt.Println(f1(3), f2(5))
}
