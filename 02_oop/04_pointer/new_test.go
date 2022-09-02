package _4_pointer

import (
	"fmt"
	"testing"
)

//new返回的是指针
func Test1(t *testing.T) {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(a)
	fmt.Println(*b) // false
	fmt.Println(b)

	*a = 100
	fmt.Println(*a)

}
