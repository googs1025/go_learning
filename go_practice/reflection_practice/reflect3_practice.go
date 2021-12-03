package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId		int
	customerId	int
}

func reflect_pr(o interface{})  {
	t := reflect.TypeOf(o)
	k := t.Kind()
	v := reflect.ValueOf(o)
	fmt.Println("Type:", t)
	fmt.Println("Kind", k)
	fmt.Println("Value:", v)

}


func main() {
	o := order{
		ordId: 1235,
		customerId: 5555,
	}

	reflect_pr(o)


}
