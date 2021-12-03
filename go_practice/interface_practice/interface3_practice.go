package main

import (
	"fmt"
)

type Person struct {
	name 	string
	sex 	string
	age 	int
}

type Student struct {
	*Person
	id 		int
	addr 	string
}

func main()  {
	s1 := Student{
		&Person{
			name: "jzy",
			sex: "man",
			age: 50001,
		},50, "hk",
	}
	fmt.Println(s1) // output:{0xc00007c4b0 50 hk}
	// 以下两个相同，会先在Student找name字段，没有才会进入Person中找!
	fmt.Println(s1.name)
	fmt.Println(s1.Person.name)

}
