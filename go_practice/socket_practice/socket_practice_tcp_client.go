package main


/*
	1.建立与服务端的连接
	2.进行数据收发
	3.关闭连接

 */


import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// 关闭连接
	defer conn.Close()
	// 建立一个存input的地方
	inputReader := bufio.NewReader(os.Stdin)
	// 循环的读数据
	for {
		input, _ := inputReader.ReadString('\n') // 用户输入数据
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {          // 如果输入Q就退出
			return
		}

		_, err = conn.Write([]byte(inputInfo))          // 发送数据
		if err != nil {
			return
		}

		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}

}
