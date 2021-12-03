package main

import "fmt"

// 结构体
type User struct {
	Name string
	Email string
}
// 定义的方法，此方法接受
// 不论是(u *User) 还是 (u User) GO会自动转换
func (u User) Notify()  {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}


func main()  {
	// 值类型调用
	u1 := User{Name: "jzy", Email: "googs1025@gmail.com"}
	u1.Notify()
	// 引用类型调用
	// 当接受者是指针时，即使用值类型调用，那函数内部也是对指针进行操作
	u2 := &User{"aaa", "aaaa@gmail.com"}
	u2.Notify()

}
