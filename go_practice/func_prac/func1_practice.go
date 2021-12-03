package main

import "fmt"

/*
	func定义函数，函数可以做为参数传递。
	建议将复杂签名定义为函数类型，方便阅读

	有返回值的函数，必须要有明确的return语句，否则会编译错误
 */

// test函数接受一个参数:fn， fn本身就是一个func() int的类型
func test(fn func() int) int {
	return fn()
}
// 将复杂函数签名定义为新的类型
type FormatFunc func(s string, x, y int) string

// 输入参数FormatFunc(一个函数)，string，int, int
func format(fn FormatFunc, s string, x ,y int) string  {
	return fn(s, x, y)

}

func main()  {
	// 输入直接是一个函数
	s1 := test(func() int {
		return 100
	})
	// 输入参数之一就是一个函数，
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s,x,y)
	}, "%d, %d",10,20)
	println(s1,s2)

}
