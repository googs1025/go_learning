package main

import (
	"fmt"
	"os"
)

/*
	向外输出函数
	Print:直接输出内容
	Println:会空行
	Printf:支持格式化输出

	Fprint:会将内容输出到一个io.Writer接口类型的变量w中
	Fprint:
	Fprintf:
	Fprintln:

	Sprint:生成输入数据，并返回字符串

	格式化占位符(常用):
	%v:值默认的格式
	%+v:与%v相同，只是如果是struct会加上字段名
	%T:打印值的类型
	%%:百分号
	%d:十进制
	%s:输入string或[]byte
 */

func main() {
	// print
	fmt.Print("aaa")
	fmt.Println("aaaa")
	name := "jzy"
	fmt.Printf("我是:%s\n", name)
	// Fprintln
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("打开文件错误m, err:", err)
		return
	}
	name1 := "jzy"

	fmt.Fprintln(fileObj, "往文件写入信息", name1)
	// Sprint
	s1 := fmt.Sprint("jzyyyyy")
	name011 := "dkjfa;ldkjf"
	age := 1555
	s2 := fmt.Sprintf("name:%s, age:%d", name011, age)
	s3 := fmt.Sprintln("kadjfkljaldfjaf")
	fmt.Println(s1,s2,s3)




}
