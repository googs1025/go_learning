package main

import "math"

/*
	匿名函数:不需要定义函数名。

 */

func main()  {
	// getSqrt为变量，是一个函数的赋值。
	getSqrt := func(a float64) (sqrt float64) {
		return math.Sqrt(a)
	}

	println(getSqrt(50))

	// 下一个阶段!! 匿名函数重要应用!!

	/*
	函数变量:fn是函数变量，可以直接调用
	 */

	fn := func() {
		println("hello world")
	}
	fn()

	/*
	切片函数:可以直接用slice存放函数
	调用的时候，用index直接调用
	 */

	fns := [](func(x int) int){
		func(x int) int { return x + 100},
		func(x int) int { return x + 500},
	}

	println(fns[0](500))

	/*
	结构体函数:可以用匿名结构体调用
	调用的时候，用.fn调用
	 */

	d := struct {
		fn func(s string) string
	}{
		fn: func(s string) string {
			return s
		},
	}

	println(d.fn("hello world"))

	/*
	通道函数:
	chan func() string
	 */

	fc := make(chan func(s string) string, 2)
	fc <- func(s string) string {
		return s
	}

	println((<-fc)("hellow world"))




}