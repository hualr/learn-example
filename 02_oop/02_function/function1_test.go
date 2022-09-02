package _2_function

import "testing"

func test1() {
	println("hello world")
}

func test2(message string) {
	println(message)
}

func test3(num int) int {
	return num * num
}

func Test1(t *testing.T) {
	test1()
}

func Test2(t *testing.T) {
	test2("hello")
}

func Test3(t *testing.T) {
	println(test3(3))
}

func test4(a, b int) int {
	return a + b
}

func Test5(t *testing.T) {
	print(test4(1, 3))
}

func test5(a, b *int) {
	var o int
	o = *a
	*a = *b
	*b = o
}

func Test6(t *testing.T) {
	a := 10
	b := 20
	test5(&a, &b)
}

func Test7(t *testing.T) {
	func(message string) {
		println(message)
	}("梨花")
}
