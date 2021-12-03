package main

import "fmt"

/*
	表达式:调用者不同，方法分为两种
	前者:method value，后者:method expression

 */
// 定义结构体
type User struct {
	id int
	name string
}
// 方法
func (self *User) Test()  {
	fmt.Printf("%p, %v\n", self, self)
}

func main()  {
	// 声明结构体
	u := User{id: 555, name: "aaaaa"}
	u.Test() // 方法调用

	//method value
	mValue := u.Test
	mValue() // 隐式传递 receiver
	//method expression
	mExpression := (*User).Test
	mExpression(&u) // 显式传递 receiver

	print("--------------------------------------")

	uu := User{id: 5555447777, name: "44"}
	mValue1 := uu.Test


	uu.id, uu.name = 222, "dddkjldfj"
	uu.Test()
	mValue1()


}