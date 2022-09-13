package channel

import (
	"fmt"
	"testing"
	"time"
)

//1. select后面的case不代表任何判断条件 只是一个信道操作
func Test11(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	for {
		select {
		case <-ch1:
			// 如果从 ch1 信道成功接收数据，则执行该分支代码
		case ch2 <- 1:
			// 如果成功向 ch2 信道成功发送数据，则执行该分支代码
		default:
			// 如果上面都没有成功，则进入 default 分支处理流程
			println("no io operation")
			time.Sleep(time.Second * 1)
		}
	}
}

func Test12(t *testing.T) {
	size := 10
	ch1 := make(chan int, size)
	for i := 0; i < size; i++ {
		ch1 <- 1
	}

	ch2 := make(chan int, size)
	for i := 0; i < size; i++ {
		ch2 <- 2
	}

	ch3 := make(chan int, 1)

	for {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case b := <-ch2:
			fmt.Println(b)
			// 因为写入一个之后就写不进去入库 因此write指挥执行一次
		case ch3 <- 10:
			fmt.Println("write")
		default:
			fmt.Println("none")
			time.Sleep(time.Second)
		}
	}
}
