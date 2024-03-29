package context

import (
	"fmt"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestC2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go f1(ctx)

	time.AfterFunc(time.Second*3, func() {
		fmt.Println("main 发出cancel信号")
		cancel()
	})

	select {
	case <-ctx.Done():
		fmt.Println("main exit")
	}
}
func f1(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	go f2(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("-----------f1 发出cancel信号-----------")
		cancel()
		fmt.Println("-----------f1 退出--------------------")
		return
	}
}

func f2(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	go f3(ctx)
	go time.AfterFunc(time.Second*2, func() {
		cancel()
	})
	select {
	case <-ctx.Done():
		fmt.Println("-----------f2 发出cancel信号-----------")
		cancel()
		fmt.Println("-----------f2 退出--------------------")
		return
	}
}

func f3(child context.Context) {
	child, cancel := context.WithCancel(child)
	go f4(child)
	select {
	case <-child.Done():
		fmt.Println("-----------f3 发出cancel信号-----------")
		cancel()
		fmt.Println("-----------f3 退出--------------------")
		return
	}
}

func f4(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	go f5(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("-----------f4 发出cancel信号-----------")
		cancel()
		fmt.Println("-----------f4 退出--------------------")
		return
	}
}

func f5(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("-----------f5 退出--------------------")
		return
	}
}
