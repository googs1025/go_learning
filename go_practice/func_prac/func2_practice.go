package main

import "fmt"

/*
	1. 值传递:调用函数时将实际参数复制一遍到函数体中，这样进行修改或计算就不会动到原变量
	2. 引用传递:定用函数时将实际参数的地址传到函数中，会改变实际参数
	3. 不管值传递还是引用传递，都是传给函数"变量的副本"。所以一般地址拷贝更为高效。值传递取决于拷贝对象的大小
	4. slice、map、channel都是引用传递

	5. 可变参数本质是slice。只有一个且是最后一个
		func myfunc(args ...int) {
		// 0个或多个参数
		}
		func add(a int, args...int) int {
		// 1个或多个参数
		}
		func add(a int, b string, args...int) }
		// 2个或多个参数
		}
		=>args本身是个slice 可以通过args[index]访问所有参数，用len(args)判断有几个参数
	6. 用interface{}传递任意类型数据，类型安全
		func myfunc(args ...interface{}) {
		}
 */


func main()  {
	var a, b = 111, 222
	println(a, b)
	/*
	调用swap()函数
	&a 指向a变量的地址
	&b 指向b变量的地址
	 */

	swap(&a, &b)
	println(a, b)
	println("========================NEXT===============================")
	println(test_new1("sum: %d", 1,5,3,67,6)) // 直接用int传入
	// 如果用slice传入做变参，必须展开 ...
	s := []int{1, 2, 5}
	println(test_new1("sum: %d", s...))


}



func test_new1(s string, n ...int) string  {
	var x int
	for _, i := range n	{
		x += i
	}
	return fmt.Sprintf(s, x)

}

func swap(x, y *int)  {
	var temp int
	// 简单地址交换!!
	temp = *x
	*x = *y
	*y = temp

}