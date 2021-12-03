package main

import "fmt"

type User struct {
	id int
	name string
}

type Manager struct {
	// 可以当做继承!!
	User // 只有类型没有名 ->匿名字段
	title string
}

// Manager 也可以用User也行继承
func (self *User) ToString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}

func (self *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v",self, self)
}


func main()  {
	// 通过匿名字段，可获得和继承类似的复用能力。
	m := Manager{User{1,"Tom"}, "boss"}
	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
	fmt.Println(m.User.ToString())


}
