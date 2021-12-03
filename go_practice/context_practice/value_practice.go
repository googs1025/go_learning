package main

import (
	"context"
	"fmt"
	"time"
)

func HandelRequset(ctx context.Context)  {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequset done")
			return
		default:
			fmt.Println("HandelRequset running, param:", ctx.Value("paramter") )
			time.Sleep(2 * time.Second)
		}
	}


}



func main() {
	// 可以在context中传递key-value
	ctx := context.WithValue(context.Background(), "paramter",222)

	go HandelRequset(ctx)
	time.Sleep(5 * time.Second)


}
