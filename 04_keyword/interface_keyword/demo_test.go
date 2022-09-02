package interface_keyword

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var a string
	//pair<statictype:string, value:"aceld">
	a = "aceld"

	//pair<type:string, value:"aceld">
	var allType interface{}
	allType = a

	str, result := allType.(string)
	fmt.Println(str, result)
}
