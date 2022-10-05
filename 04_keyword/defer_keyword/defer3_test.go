package defer_keyword

import (
	"fmt"
	"testing"
)

func Test3(t *testing.T) {
	defer func() {
		if ok := recover(); ok != nil {
			fmt.Println("recover")
		}
	}()

	panic("error")
}

func ShowDefer() {
	fmt.Println("最后输出值:", deferValue())
}
func deferValue() (ret int) { // 注意这里返回res变量值
	ret = 0

	defer func() { // 会直接修改栈中对应的返回值
		println(ret)
		ret += 10
		fmt.Println("Defer 运行值:", ret)
	}()
	ret = 2
	fmt.Println("Ret重置值:", ret)
	return //返回的ret最后是 其实是本次2+上面的ret+=10的结果
}

func Test4(t *testing.T) {
	ShowDefer()
}

func ShowDeferA() {
	fmt.Println("最后输出值:", deferValueA())
}
func deferValueA() int { // 非命名变量返回
	ret := 0
	defer func() {
		ret += 10
		fmt.Println("Defer 运行值:", ret)
	}()
	ret = 2
	return ret // 这里直接返回ret2
}

func Test5A(t *testing.T) {
	ShowDeferA()
}
