package main

import "fmt"

/*
	接口:定义对象的行为规范，只有定义规范不具体实现。由对象本身或方法来实现具体细节
	go中的接口为一种类型，是一组method集合。接口像是定义一种协议，不关心数据属性，只有关心行为。

	接口是一个或多个方法签名的集合
	任何类型的方法集中只要拥有该接口"对应的全部方法"签名 ->表示实现该接口
	接口只有方法声明，没有实现，没有数据字段

	type 接口类型名 interface {
		方法名1 (param_list) return_list
		方法名2 (param_list) return_list
		...
	}

	例如:
	type writer interface {
		Write( []byte) error
	}

 */

func main()  {
	// 没有实现接口的方法
	//c := Cat{}
	//fmt.Println("猫:", c.Say())
	//d := Dog{}
	//fmt.Println("狗:", d.Say())


	// 声明一个Sayer接口变量x
	// 接口变量可以存储所有实现该接口的实例!!
	// 所以Sayer类型变量可以存储dog和cat类型的变量
	var x Sayer
	// 实例化
	a := Cat{}
	b := Dog{}
	// 把实例赋值给x
	x = a
	x.Say()
	x = b
	x.Say()
}

type Cat struct {

}
// 所以cat对象 实现了Sayer接口，因为它将Sayer接口中的方法全部都实现了!!
func (cat Cat) Say()  {
	fmt.Println("喵喵喵!")
}
// dog对象，也实现了Sayer接口
type Dog struct {

}

func (dog Dog) Say()  {
	fmt.Println("旺旺旺!")
}
// Sayer接口里只有一个Say方法，所以只需要给dog和cat分别实现say方法即可
type Sayer interface {
	Say()
}

