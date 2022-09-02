package main

import (
	"fmt"
	"testing"
)

func Test4(t *testing.T) {
	s := []int{1, 2, 3} //len = 3, cap = 3, [1,2,3]

	//1. [0, 2) 注意边界
	s1 := s[0:2] // [1, 2]

	fmt.Println(s1)
	s1[0] = 100

	//2. 这两切片是同源的 也就是说一个改了 另外一个也会改
	fmt.Println("s=", s, "s1=", s1)

	//3. copy 可以将底层数组的slice一起进行拷贝 修改不同源
	s2 := make([]int, 3) //s2 = [0,0,0]

	//将s中的值 依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2)
	s[0] = 1000
	fmt.Println(s2, s)
}
