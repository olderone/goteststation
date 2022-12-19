package main

import (
	"context"
	"fmt"
	"time"
)

func MyOperate1(ctx context.Context) {
	for {
		select {
		default:
			fmt.Println("MyOperate1", time.Now().Format("2006-01-02 15:04:05"))
			time.Sleep(2 * time.Second)
		case <-ctx.Done():
			fmt.Println("MyOperate1 Done")
			return
		}
	}
}
func MyOperate2(ctx context.Context) {
	fmt.Println("Myoperate2")
}
func MyDo2(ctx context.Context) {
	go MyOperate1(ctx)
	go MyOperate2(ctx)
	for {
		select {
		default:
			fmt.Println("MyDo2 : ", time.Now().Format("2006-01-02 15:04:05"))
			time.Sleep(2 * time.Second)
		case <-ctx.Done():
			fmt.Println("MyDo2 Done")
			return
		}
	}

}
func MyDo1(ctx context.Context) {
	go MyDo2(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("MyDo1 Done")
			// 打印 ctx 关闭原因
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("MyDo1 : ", time.Now().Format("2006-01-02 15:04:05"))
			time.Sleep(2 * time.Second)
		}
	}
}
func main() {
	// 创建 cancelCtx 实例
	// 传入context.Background() 作为根节点
	ctx, cancel := context.WithCancel(context.Background())
	// 向协程中传递ctx
	go MyDo1(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("stop all goroutines")
	// 执行cancel操作
	cancel()
	time.Sleep(2 * time.Second)
}
