package main

import (
	"fmt"
	"reflect"
)
/*
	静态类型: ex:int、float32 []byte :每个变量都有个静态变量，在编译时就确定

	interface:特殊类型，方法集合

	反射:检查interface的(value, type)
	反射就是程序能够在运行时检查变量和值，求出它们的类型。

	reflect.Type:提供一组接口处理interface类型，(value,type)的type
	reflect.Value:提供一组接口处理interface类型，(value,type)的value

 */


func main() {
	var x float64 = 3.22
	t := reflect.TypeOf(x)
	fmt.Println("type:", t)

	v := reflect.ValueOf(x)
	fmt.Println("value:", v)


	var y float64 = v.Interface().(float64)
	fmt.Println("value:", y)
}
