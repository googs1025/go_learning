package main

import (
	"fmt"
	"net/http"
)

/*
	wed服务器可以归纳为:
	1.客户端通过tcp/ip协议建立到服务端的tcp连接
	2.客户端向服务器发送http请求包，请求服务器的资源文挡
	3.服务器向客户端发送http响应包
	4.客户端与服务端断开。由客户端解释html文挡，在客户端上渲染图形效果
 */

func main() {
	// 单独写回调函数
	http.HandleFunc("/go", myHander)
	// addr: 监听地址， hander: 回调函数
	http.ListenAndServe("127.0.0.1:8000",nil)

}
// hander函数
func myHander(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式: GET POST DELETE PUT UPDATE
	fmt.Println("method:", r.Method)
	// /go
	fmt.Println("url:", r.URL.Path)

	fmt.Println("header:", r.Header)

	fmt.Println("body:", r.Body)
	// 回复
	w.Write([]byte("www.51mh.com"))

}
