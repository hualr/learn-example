package channel

import (
	"fmt"
	"testing"
)

func SelectForChan(c chan string) {
	var rev string
	send := "Hello"
	select {
	case rev = <-c:
		fmt.Printf("recvied %s \n", rev)
	case c <- send:
		fmt.Printf("sent %s \n", send)
	}
}
func SelectForChan2(c chan string) {
	var rev string
	send := "Hello"
	select {
	case rev = <-c:
		fmt.Printf("recvied %s \n", rev)
	case c <- send:
		fmt.Printf("sent %s \n", send)
	default:
		fmt.Printf("no data found in default \n")
	}
}

// 由于没有缓冲 因此此时直接报错 select不能读写
func Test21(t *testing.T) {
	c := make(chan string)
	SelectForChan(c)
}

// 这里的default避免了deadlock问题
func Test21_2(t *testing.T) {
	c := make(chan string)
	SelectForChan2(c)
}

/**
select for chan 是每次尝试读或者写
首先我们是空 因此没法读 只能写
当我们尝试写之后 select监听到写的操作 因此可以匹配到case2 send
*/
func Test22(t *testing.T) {
	c := make(chan string, 1)
	SelectForChan(c)
}

/**
此时已经可以读 因此监听到读的操作
*/
func Test23(t *testing.T) {
	c := make(chan string, 1)
	c <- "hello"
	SelectForChan(c)
}

/**
可以读或者写
*/
func Test24(t *testing.T) {
	c := make(chan string, 2)
	c <- "hello"
	SelectForChan(c)
}
