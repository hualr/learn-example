package defer_keyword

import "testing"

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func Test5(t *testing.T) {
	println(f1())
	println(f2())
	println(f3())

	/**
	改写为 返回值=xxx
	调用defer
	空的return
	*/
}
