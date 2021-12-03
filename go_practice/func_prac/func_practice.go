package main

import "fmt"

/*
	golang函数特点:
	1.不需声明原型
	2.支持不定变参
	3.支持多返回值
	4.支持匿名函数与闭包
	5.函数也是一种类型，一个函数可以赋值给一个变量

	*不能嵌套 一个包不能有两个名一样的函数
	*不能重载
	*不能有默认参数
 */


func main()  {
	a,_ := test1(1,2, "aaaa")
	fmt.Println(a)
	b := test2(555,666)
	fmt.Println(b)


}
// 范例一
func test1(x, y int, s string) (int, string) {
	// 类型相同的相邻参数，参数类型可以合并，多返回值需要括号。
	n := x + y
	return n, fmt.Sprintf(s, n)

}

func test2(x, y int) (m int) {
	// 返回值可以直接在签名处赋值，然后直接用即可，不需要再命名
	m = x + y
	return

}

