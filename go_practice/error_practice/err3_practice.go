package main

/*
	自定义错误
 */

import (
	"fmt"
	"os"
	"time"
)
// 定义struct错误:
type PathError struct {
	path 		string
	op 			string
	createTime 	string
	message 	string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s",
		p.path, p.op, p.createTime, p.message)
}

func Open(filename string) error {

	file,err := os.Open(filename)
	//错误处理
	if err != nil {
		return &PathError{
			path: filename,
			op: "read",
			createTime: fmt.Sprintf("%v", time.Now()),
			message: err.Error(),

		}
	}

	// file.Close也处理!!
	if file != nil {
		defer func() {
			if err := file.Close(); err != nil {
				fmt.Printf("defer close txt err %v\n", err)
			}
		}()
	}


	// 业务逻辑!!!

	return nil
 }

func main()  {
	err := Open("test.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error", v)
	default:
		return


	}

}
