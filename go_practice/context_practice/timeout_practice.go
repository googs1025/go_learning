package main

import (
	"context"
	"fmt"
	"time"
)

func HandelRequset(ctx context.Context)  {
	go WriteDatabase(ctx)
	go WriteRedis(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequset done")
			return
		default:
			fmt.Println("HandelRequset running")
			time.Sleep(2 * time.Second)
		}
	}


}

func WriteRedis(ctx context.Context)  {
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
	// 设定时间退出
	ctx, _:= context.WithTimeout(context.Background(), 5 * time.Second)

	go HandelRequset(ctx)




	time.Sleep(10 * time.Second)
}

