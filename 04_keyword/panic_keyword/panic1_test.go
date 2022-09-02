package panic_keyword

import (
	"fmt"
	"testing"
)

func f1() {
	fmt.Println("func f1 start")
	var arr []int
	fmt.Println(arr[10])
	fmt.Println("func f1 end")
}

func Test1(t *testing.T) {
	f1()
}

// recover没有机会执行
func Test2(t *testing.T) {
	f1()
	r := recover()
	fmt.Printf("%s \n", r)
	fmt.Println("main func end")
}

func Test3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s\n", r)
		}
		fmt.Println("main func end")

	}()
	f1()
}

func Test4(t *testing.T) {
	panic(" init mysql failed")
}
