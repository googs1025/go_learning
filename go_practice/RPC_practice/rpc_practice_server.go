package main

import (
	"log"
	"net/http"
	"net/rpc"
)

/*
	RPC 远程调用 -> 一种通信协议
	协议允许运行于一台计算机的程序调用另一台计算机的子程序，且两者不需额外交互

	场景:实现rpc通信
 */


// 参数对像
type Params struct {
	Width int
	Height int
}
// 用于注册用!!
type Rect struct {

}
// RPC方法
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func (r *Rect) Area(p Params, ret *int) error  {
	*ret = p.Height * p.Width
	return nil
}

func main() {
	// 1.new一个注册服务
	rect := new(Rect)
	// 注册一个服务 rect
	rpc.Register(rect)
	// 2.服务处理绑定到http协议上
	rpc.HandleHTTP()
	// 3.监听服务
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Panicln(err)
	}
}
