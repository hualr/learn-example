package interface_keyword

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var a string
	a = "aceld"

	var allType interface{}
	allType = a

	str, result := allType.(string)
	fmt.Println(str, result)
}

type dog struct {
}

func Test2(t *testing.T) {
	var a interface{}
	a = dog{}
	switch t := a.(type) {
	default:
		fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	case dog:
		fmt.Printf("dog 类型")
	}
}
