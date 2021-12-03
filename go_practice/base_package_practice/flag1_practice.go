package main

import (
	"fmt"
	"os"
)

/*
	使用flag包对命令行解析:
	os.Args:想获得命令行参数
	os.Args是一个[]string类型

 */

func main() {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n",index, arg)
		}
	}
	//go run flag1_practice.go
	//返回值结果
	//args[0]=C:\Users\dodo\AppData\Local\Temp\go-build252196574\b001\exe\flag_practice.exe
	//args[1]=dfadf
	//args[2]=adkfj
	//args[3]=fakdjfladjf
	//args[4]=243



}
