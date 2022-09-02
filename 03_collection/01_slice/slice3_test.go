package main

/**
len描述的是slice本身的长度
cap表示的是slice当前的容量
cap满了可以自动double扩容
*/
import (
	"fmt"
	"testing"
)

func Test3(t *testing.T) {
	numbers := make([]int, 3, 4)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素1, numbers len = 4， [0,0,0,1], cap = 5
	numbers = append(numbers, 1)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向一个容量cap已经满的slice 追加元素，自动扩容了
	numbers = append(numbers, 2)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

}
