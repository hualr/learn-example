package _4_defer

import (
	"errors"
	"fmt"
	"testing"
)

// defer的基本使用 多个defer执行顺序相反
func Test1(t *testing.T) {
	//写入defer关键字
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}

// defer处理异常
func Test2(t *testing.T) {
	fmt.Println("Enter function main.")
	defer func() {
		fmt.Println("Enter defer function.")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer function.")
	}()
	// 引发 panic。
	panic(errors.New("something wrong"))
	fmt.Println("Exit function main.")
}

func main() {

}
