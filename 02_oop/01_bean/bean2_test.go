package _1_bean

import (
	"fmt"
	"testing"
)

type cat struct {
	age int
}

// 一个返回的是指针 一个不是指针
func newCat1() *cat {
	return &cat{}
}

func newCat2() cat {
	return cat{}
}

//接收器和函数的区别在于 函数没有类型 接收器有
func (c cat) eat1() {
	fmt.Println("cat is eating")
}

func (c *cat) eat2() {
	fmt.Println("cat is eating")
}

//接收器和函数的区别在于 函数没有类型 接收器有
func (c cat) growUp1(age int) {
	c.age = age
}

func (c *cat) growUp2(age int) {
	c.age = age
}

func Test4(t *testing.T) {

	c2 := newCat1()
	c3 := newCat2()

	//接收器的类型不管是值还是指针 都可以被指针或者值调用
	c2.eat1()
	c3.eat1()

	c2.eat2()
	c3.eat2()

	c2.growUp1(1)
	println(c2.age)

	c2.growUp2(1)
	println(c2.age)

	c3.growUp1(1)
	println(c3.age)

	c3.growUp2(1)
	println(c3.age)
}
