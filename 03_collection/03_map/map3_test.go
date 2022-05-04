package main

import "testing"

func Test3(t *testing.T) {
	// 1.取出来不一定是真的存在
	m := make(map[string]string)
	s, ok := m["one"]
	if ok {
		return
	}
	print(s)

	// 2.不初始化 就无法进行增 但是可以进行其他操作
	var m2 map[string]int

	print(m2)
	print(m)

}
