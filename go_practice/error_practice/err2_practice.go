package main

import (
	"errors"
	"fmt"
)
/*
	异常对象返回
 */
func getCircleArea(redius float32) (area float32, err error) {
	if redius < 0 {
		// 构建异常对象
		err = errors.New("半径不能为负")
		return
	}
	area = 3.14*redius*redius
	return
}

func main()  {
	area, err := getCircleArea(-5)
	// 抛异常
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("area is :", area)
	}

}
