package _go

import (
	"fmt"
	"testing"
	"time"
)

func running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)
		// 延时1秒
		time.Sleep(time.Second)
	}
}

func Test1(t *testing.T) {
	// 并发执行程序
	go running()

	for {
		time.Sleep(1)
	}
}

func Test2(t *testing.T) {
	go func(start int) {
		var times int
		for {
			times++
			fmt.Println("tick", times)
			time.Sleep(time.Second)
		}
	}(1)

	for {
		time.Sleep(1)
	}
}
