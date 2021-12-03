package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	读写互斥锁
	读多写少场景!
	在读操作访问时的时后不用限制其他goroutine
	在写操作访问时需要限制goroutine访问

	模拟需求:读者与写者并发问题
 */

var (
	x2 		int
	wg2 	sync.WaitGroup
	//mu2 	sync.Mutex
	rwlock 	sync.RWMutex
)

func read()  {
	//mu2.Lock()
	rwlock.RLock()
	time.Sleep(time.Microsecond)
	rwlock.RUnlock()
	//mu2.Unlock()
	wg2.Done()
}

func write()  {
	//mu2.Lock()
	rwlock.Lock()
	x2 += 1
	time.Sleep(10 * time.Microsecond)
	rwlock.Unlock()
	//mu2.Unlock()
	wg2.Done()


}



func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go write()
	}


	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go read()
	}

	wg2.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))

	
}
