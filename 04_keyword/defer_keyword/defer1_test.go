package defer_keyword

import (
	"fmt"
	"testing"
)

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

//1 defer的值在defer的时候已经被确定
func Test1(t *testing.T) {
	a()
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

// 在return之后执行
func Test2(t *testing.T) {
	println(c())
}
