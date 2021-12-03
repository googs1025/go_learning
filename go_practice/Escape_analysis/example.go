package main

type Student struct {
	Name 	string
	Age 	int
}
// GO 可以返回局部变量的指针类型 *Student
func StudentRegister(name string, age int) *Student {

	s := new(Student)

	s.Name = name
	s.Age = age
	return s
}

//func main() {
//	StudentRegister("jim", 10)
//	//fmt.Println(a)
//}