package _2_ken

import (
	"fmt"
	"reflect"
	"testing"
)

func TestD1(t *testing.T) {
	var s *string

	var i interface{} = s

	if i == nil {
		fmt.Println("it's empty")
	} else {
		fmt.Println("it's non-empty", i)
	}

	if reflect.ValueOf(s).IsNil() {
		fmt.Println("it's empty")
	} else {
		fmt.Println("it's non-empty")
	}
}
