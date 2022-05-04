package main

// 孔乙己
import (
	"fmt"
	"testing"
)

func Test2(t *testing.T) {
	//1. 普通写法
	slice1 := []int{1, 2, 3}
	fmt.Println(len(slice1), slice1)

	//2. 声明slice1是一个切片，但是并没有给slice分配空间
	var slice2 []int
	fmt.Println(len(slice2), slice2)

	//3. 开辟3个空间 ，默认值是0 这个是值得被推广的
	slice3 := make([]int, 3)
	fmt.Println(len(slice3), slice3)

}
