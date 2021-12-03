package main

import "fmt"

/*
	值接收者 vs 指针接收者
	值接收者实现接口后，不论是T结构体还是*T结构体指针都可以赋值给接口变量 (内部会自动转换)
	指针接收者则不能将T类型传给x，只能存储*T
 */

type Mover interface {
	Move()
}

type Dog1 struct {

}

func (dog Dog1) Move()  {
	fmt.Println("狗动了")
}
//指针接收者
//func (dog *Dog1) Move()  {
//	fmt.Println("狗动了")
//}

func main()  {
	var x Mover
	var jzy = Dog1{}
	x = jzy
	x.Move()
	var jjj = &Dog1{}
	x = jjj
	x.Move()
}