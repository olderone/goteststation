package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		wg sync.WaitGroup
	)
	ctx := context.Background()
	cancelCtx, cancel := context.WithTimeout(ctx, time.Second*9)
	defer cancel()
	wg.Add(1)
	go task1(cancelCtx, &wg)
	wg.Wait()
}

func task1(ctx context.Context, wg *sync.WaitGroup) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err()) // context deadline exceeded
			wg.Done()
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}
