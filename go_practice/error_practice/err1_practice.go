package main

import "fmt"

/*
	系统抛出异常
 */

//func test01()  {
//	a := [5]int{0,1,2,3,4,5}
//	a[1] = 123
//	fmt.Println(a)
//
//	index := 10
//	a[index] = 10
//	fmt.Println(a)
//
//}

func getCircleArea(redius float32) (area float32)  {
	if redius < 0 {
		// 自己抛
		panic("半径不能为负")
	}
	return 3.14*redius*redius

}

func test03()  {
	// 延时执行匿名函数
	// 延时到何时? 1. 程序正常结束 2. 发生异常时

	defer func() {
		// recever() 复活 恢复
		// 返回程序为啥挂了
		if err := recover();err != nil {
			fmt.Println(err)
		}
	}()

	getCircleArea(-5)
	fmt.Println("这里没有执行")

}

func test04()  {
	test03()
	fmt.Println("test04")

}

func main()  {
	test04()


}
