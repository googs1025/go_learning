package main

import (
	"fmt"
	"reflect"
)

/*
	可以透过反射，将原本通用性不高的代码转变为通用姓高的
	采用interface{}!!
 */
 */

type order struct {
	ordId		int
	customerId	int
}

type employee struct {
	name	string
	id 		int
	address string
	salary	int
	country	string
}


func create_reflect(o interface{})  {
	// 判断是不是struct类型
	if reflect.ValueOf(o).Kind() == reflect.Struct {

		t := reflect.TypeOf(o).Name()
		query := fmt.Sprintf("insert into %s values",t)
		v := reflect.ValueOf(o)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s\"%s\"",query, v.Field(i).String())
				}
			default:
				// 如果struct中有字段不是int or string
				fmt.Println("Unsurpported type")
				return


			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")

}


func main() {
	o := order{
		ordId: 123,
		customerId: 455,
	}
	create_reflect(o)
	e := employee{
		name: "jzy",
		id: 123,
		address: "googs1025@gmail.com",
		salary: 456,
		country: "bj",
	}
	create_reflect(e)
	i := "staaaaaa"
	create_reflect(i)

}
