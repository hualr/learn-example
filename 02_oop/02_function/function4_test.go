package _2_function

import "testing"

//落回报
func test6(q, p int) (a int, b int) {
	a = q + p
	b = q - p
	return
}

func Test9(t *testing.T) {
	print(test6(1, 8))
}
