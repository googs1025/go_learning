package main

import (
	"flag"
	"fmt"
	"time"
)

/*
	定义命令行flag参数
	flag.Type()
	flag.TypeVar()
 */

func main() {
	var (
		name 	string
		age 	int
		married bool
		delay 	time.Duration
	)
	// 写入要输入参数
	flag.StringVar(&name, "name", "张三","姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d",0,"延迟时间间隔")
	// 进行解析
	flag.Parse()
	fmt.Println(name, age, married, delay)
	// 返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	// 返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	// 返回命令行参数后的参数个数
	fmt.Println(flag.NFlag())

}
