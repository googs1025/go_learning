package main

import "fmt"

type order struct {
	ordId 		int
	customerId	int
}


func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}


func main() {
	o := order{
		ordId: 1234,
		customerId: 55555,
	}
	a := createQuery(o)
	fmt.Println(a)
}
