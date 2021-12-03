package main

/*
	给定一个结构体S和结构体T的类型，如:类型S包含类型T，则S和*S方法集包含T方法
 */

import "fmt"

type S struct {
	T
}

type T struct {
	int
}

func (t T) testT()  {
	fmt.Println("如果类型 S 包含匿名字段 T，则 S 和 *S 方法集包含T方法")
}

func (t *T) testP()  {
	fmt.Println("如果类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含*T方法")
}

func main()  {
	s1 := S{T{555}}
	s2 := &S{T{555555444}}
	fmt.Printf("s1 is : %v\n", s1)
	s1.testT()
	s1.testP()
	fmt.Printf("s2 is : %v\n", s2)
	s2.testT()
	s2.testP()

}
