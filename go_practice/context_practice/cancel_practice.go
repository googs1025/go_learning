package main



import (
	"context"
	"fmt"
	"time"
)

/*
	context使用场景:可以控制一组goroutine 使得每组goroutine都有上下文
	context接口定义四种方法:
	Deadline()
	Done():需要返回一个chan
	Err():需要返回一个error
	Value()

	context使用方法:主要是三个
	*:context.WithCancel:
	1.需要初始化ctx实例
	2.将cancel()实例添加到父节点的children中
	3.返回ctx实例与cancel()方法
	4.并且在需要cancel时，调用ctx.Done() 告知下游goroutine需要退出的逻辑

	*:context.WithTimeout
	基本与context.WithCancel相同，区别在需要多设置一个到期时间参数。

	*:context.WithValue
	需要设置 key-value参数 并且可以用 ctx.Value(key) 在下游goroutine调用


	场景: 需要对database写入，
	在main goroutine调用HandelRequset，并且设置context
	并且在HandelRequest中启两个goroutine


 */

func HandelRequest(ctx context.Context)  {
	// 启两个!! 分别做两件事
	go WriteDatabase(ctx)
	go WriteRedis(ctx)
	// 不断检查是否需要cancel!!
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest done")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(2 * time.Second)
		}
	}


}

func WriteRedis(ctx context.Context)  {
	// 检查是否需要cancel
	for {
		select {
		case <-ctx.Done():
			fmt.Println("write redis done")
			return
		default:
			fmt.Println("writing redis!!")
			time.Sleep(3 * time.Second)
		}
	}

}

func WriteDatabase(ctx context.Context)  {
	// 检查是否需要cancel
	for {
		select {
		case <-ctx.Done():
			fmt.Println("write database done")
			return
		default:
			fmt.Println("writing database")
			time.Sleep(2 * time.Second)


		}
	}
}

func main() {
	// 在main gorourine中需要先调用ctx!!
	ctx, cancel := context.WithCancel(context.Background())
	// 启一个goroutine去处理
	go HandelRequest(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("it is time to stop all sub goroutine")
	// 需要取消
	cancel()
	fmt.Println("finish the cancel!")
	time.Sleep(3 * time.Second)
}
