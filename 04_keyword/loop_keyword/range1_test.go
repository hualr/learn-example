package loop_keyword

import (
	"fmt"
	"testing"
)

/**
1. 关注怎么写的range
2. 这里number1 |的处理需要值得注意
*/
func Test1(t *testing.T) {
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	fmt.Println(numbers1)
}

// 这里关注 range可以拿到index和value
func Test2(t *testing.T) {
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for index, value := range numbers2 {
		if index == maxIndex2 {
			numbers2[0] += value
		} else {
			numbers2[index+1] += value
		}
	}
	fmt.Println(numbers2)
}
