package main

import "fmt"

/*
	方法集:每个类型都有与之关联的方法集
	1. 类型T 方法集包含全部 receiver T方法
	2. 类型T *T方法集包含全部 receiver T + *T
	3. 如类型 S 包含匿名字段 T，则S和*S 方法集包含 T方法
	4. 如类型 S 包含匿名字段 *T，则S和*S 方法集包含 T和*T方法
	5. 不管嵌入 T或*T, *S方法集总包含 T + *T 方法
 */


type T struct {
	int
}

func (t T) testT()  {
	fmt.Println("类型 T 方法集包含全部 receiver T 方法")
}

func (t *T) testP()  {
	fmt.Println("类型 *T 方法集包含全部 receiver *T 方法")
}

func main()  {
	t1 := T{1}
	fmt.Printf("t1 is : %v\n", t1)
	t1.testT()
	t1.testP()

	t2 := &T{5}
	fmt.Printf("t2 is : %v\n", t2)
	t2.testP()
	t2.testT()

}
