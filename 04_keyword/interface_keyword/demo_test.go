package interface_keyword

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var a string
	a = "aceld"

	//1. interface是所有结构体的父类
	var allType interface{}
	allType = a
	//2. 断言判断一个类是否为string
	str, result := allType.(string)
	fmt.Println(str, result)
}

type dog struct {
}

func Test2(t *testing.T) {
	var a interface{}
	a = dog{}
	// type关键词主要用于和switch结合判断
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

// 空接口存在两个值 一个是类型 一个是值
func Test3(t *testing.T) {
	var i interface{}
	fmt.Printf("type: %T, value: %v", i, i)
}

func myfunc(iface interface{}) {
	fmt.Println(iface)
}

func myfunc2(ifaces ...interface{}) {
	for _, iface := range ifaces {
		fmt.Println(iface)
	}
}

// 空接口常常用来接收任何值
func Test4(t *testing.T) {
	a := 10
	b := "hello"
	c := true

	myfunc(a)
	myfunc(b)
	myfunc(c)

	myfunc2(a, b, c)
}

// 反向赋值不可取
func Test5(t *testing.T) {
	//var a int = 1
	//
	//var i interface{} = a
	//
	//var b int = i
}

// 其实这两主要是在函数中容易犯错
func Test6(t *testing.T) {
	//sli := []int{2, 3, 5, 7, 11, 13}
	//
	//var i interface{}
	//i = sli
	//
	//g := i[1:3]
	//fmt.Println(g)
}
