package main

import (
	"fmt"
	"log"
	"os"
)

/*
	log实现日志服务

	log配置
	log.SetFlags

	log前缀
	log.SetPrefix

	创建logger
	log.New

	配置log输出位置
	log.SetOutput

 */

func init()  {
	logFile, err := os.OpenFile("./logger.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:",err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)

}

func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一个普通的日志")
	log.SetPrefix("[pprof]")
	log.Println("这也是一个普通的日志")

	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是一个自定义logger记录的日志")
	logFile, err := os.OpenFile("./logger.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:",err)
		return
	}
	logger.SetOutput(logFile)




}
