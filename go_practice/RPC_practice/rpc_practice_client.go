package main

import (
	"fmt"
	"log"
	"net/rpc"
)

/*
	写RPC程序时，必须符合四个条件
	1.结构体字段需要大写! 别人需要调用
	2.函数名需要大写! 别人需要调用
	3.函数接受的第一个参数是接收参数，第二个参数是返回给客户端的参数，必为指针
	4.函数必须返回一个error类型
 */


type Params struct {
	Width 	int
	Height 	int
}


func main() {

	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	ret := 0
	err2 := conn.Call("Rect.Area", Params{500, 1000}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("面积:",ret)

	err3 := conn.Call("Rect.Perimeter", Params{11111, 5000}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("周长:",ret)
}
