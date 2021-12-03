package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
	go内置map不是并发安全。
	法一:在少量时可以运行

	sync.Map 是并发安全
	并且提供 Store Load等方法

 */

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int)  {
	m[key] = value
}


var m0 = sync.Map{}


func main() {
	//////////////////////////////////////////////////////
	//wg := sync.WaitGroup{}
	//
	//for i := 0; i < 7; i++ {
	//	wg.Add(1)
	//
	//	go func(n int) {
	//		key := strconv.Itoa(n)
	//		set(key, n)
	//		fmt.Printf("k=:%v,v:=%v\n", key, get(key))
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()
	//
	/////////////////////////////////////////////////////////
	wg1 := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg1.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m0.Store(key, n)
			value, _ := m0.Load(key)
			fmt.Printf("k=%v, v:=%v\n", key, value)
			wg1.Done()
		}(i)
	}
	wg1.Wait()


}
