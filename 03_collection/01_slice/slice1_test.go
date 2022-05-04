package main

/**
1. 理解切片和数组的写法区别
2. 理解切片的值 在函数中是否可以改变
*/
import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	// 1. 切片和数组的区别在于 切片没有指定数组的大小
	slice1 := []int{1, 2, 3, 4} // 动态数组，切片 slice

	fmt.Printf("slice1 type is %T\n", slice1)
	//2. 传入本身 就可以改变本身 而不需要传入引用
	printArray(slice1)

	fmt.Println(" ==== ")

	printArray(slice1)
}

func printArray(myArray []int) {
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100
}
