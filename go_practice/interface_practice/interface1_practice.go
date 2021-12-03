package main

import "fmt"

/*
	仅提供类型而不写字段名。
 */


type Person struct {
	name string
	sex string
	age int
}


type Student struct {
	Person // 匿名字段 (可以类似继承)
	id int
	addr string
	// 在不同结构体中，同名字段
	name string
}

func main()  {

	s1 := Student{
		Person{
			name: "jzy",
			sex: "man",
			age: 20,
		},1,"bj", "adkfjlsdjf"}
	fmt.Println(s1)

	s2 := Student{Person: Person{"aaa", "woman", 50}}
	fmt.Println(s2)

	s3 := Student{Person: Person{name: "aaaaaaa"}}
	fmt.Println(s3)
	// 对已经实例化的结构体字段改名
	s1.Person.name = "dkfjlksdjfl"
	s1.name = "bbbbbbb"

	fmt.Println(s1)
}