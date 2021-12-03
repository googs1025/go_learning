package main

/*
	go方法一定要绑定对象实例，并且将实例作为第一实参(receiver)
	1.只能为当前包内命名类型定义方法
	2.receiver参数类型可以是T或是*T类型，不能是指针与接口
	3.可以用实例的value或pointer调用全部方法

	方法定义:
	func (receiver type) methodName(para_list) (return_list) {
	}
 */

// test是一个空对象
type test struct {

}
// method1是接受test实例的方法，接受单一参数
func (t test) method1(i int) {

}
// method2 多参数，无返回值
func (t test) method2(x, y int)  {

}
// 无参数，单一返回值
func (t test) method3() (a int)  {
	return
}
// 接受两个参数，返回两个参数
func (t test) method4(x, y int) (z int, err error)  {
	return
}
// 无输入参数，无返回参数
func (t *test) method5()  {

}
// 单参数，无返回值
func (t *test) method6(i int)  {

}
// 多参数，无返回值
func (t *test) method7(x, y int)  {
}

// 多参数，单返回值
func (t *test) method8(x, y int) (z int, err error) {
	return
}

// 多参数，多返回值
func (t *test) method9(x, y int) (z int, err error) {
	return
}

func main()  {

}




