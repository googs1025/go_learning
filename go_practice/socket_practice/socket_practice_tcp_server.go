package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
	socket是unix的进程通信机制，称为套接字 -> 描述ip地址+端口号 -> TCP/IP网络地址的API
	常用有两种格式: 流式socket和数据报文式socket
	流式:面向连接的socket 针对于面向连接的tcp
	数据报文式:面向无连接socket 针对于面向无连接的udp
 */

/*
	tcp/ip协议 面向连接、可靠、基于字节流的传输层通信协议。
	数据向流水一样传输，会出现粘包问题

	tcp服务端
	一个tcp服务端可以同时连接多个client 在go中采用多个goroutine并发
	可以建立一次连接就起一个goroutine

	处理流程:
	1.监听端口
	2.接收客户端请求连接
	3.创建goroutine处理连接

 */

func process(conn net.Conn)  {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Println("收到client端的书据:", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据

	}
}


func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("lsiten failed err:", err)
		return
	}

	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		go process(conn)			// 启动一个goroutine处理连接
	}


}
