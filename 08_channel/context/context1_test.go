package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func contextTest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// 被取消或者超时就结束协程
			println("goroutin finished")
			return
		default:
		}
		// 每隔 1 秒钟，打印 running
		time.Sleep(time.Second)
		println("running")
	}
}

func Test1(t *testing.T) {
	// 返回值 cancelFunc 是一个函数，用于取消运行中的协程
	ctx, cancelFunc := context.WithCancel(context.Background())
	go contextTest(ctx)

	// 让contextTest 协程运行 3 秒钟，然后调用取消函数
	time.Sleep(3 * time.Second)
	println("send a closed signal")
	cancelFunc()

	// 等待 3 秒钟，让 contextTest 协程优雅结束。
	time.Sleep(3 * time.Second)
}

func Test2(t *testing.T) {
	// 3 秒后自动取消运行中的协程
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	go contextTest(ctx)

	// 等待 5 秒钟，让 contextTest 协程优雅结束。
	time.Sleep(5 * time.Second)
}

func Test3(t *testing.T) {
	// 3 秒后自动取消运行中的协程
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	go contextTest(ctx)

	// 等待 5 秒钟，让 contextTest 协程优雅结束。
	time.Sleep(5 * time.Second)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf("处理响应 用户名:%v 密码:%v",
		ctx.Value("UserName"),
		ctx.Value("PassWord"),
	)
}

func Test4(t *testing.T) {
	UserName := "dd"
	PassWord := "12345"
	ctx := context.WithValue(context.Background(), "UserName", UserName)
	ctx = context.WithValue(ctx, "PassWord", PassWord)
	HandleResponse(ctx)
}
