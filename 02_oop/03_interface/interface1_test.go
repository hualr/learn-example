package _3_interface

import (
	"fmt"
	"testing"
)

type action interface {
	talk()
	reName(message string)
}

type teacher struct {
	name string
}

//当实现了一个接收者是值类型的方法，就可以自动生成一个接收者是对应指针类型的方法
func (t teacher) talk() {
	println("class begin")
}

//接收器其实没什么关系 关键是是否需要修改或者是否比较大
func (t *teacher) reName(name string) {
	t.name = name
}

func Test1(t *testing.T) {
	//注意到接口接收者
	//其实,只有这个对象的指针才实现了这个接口,而对象本身其实没有实现这个接口
	/*
		var t1 action= teacher{}
		t1.reName("王平")
		t1.talk()
		fmt.Println(t1)
	*/

	var t1 teacher = teacher{}
	t1.reName("王平")
	t1.talk()
	fmt.Println(t1)

	var t2 action = &teacher{}
	t2.reName("李茂")
	t2.talk()
}

// 判断对象类型的方法
func Test2(t *testing.T) {
	var t1 teacher = teacher{}
	_, ok := interface{}(t1).(action)
	fmt.Printf("result: %v\n", ok)

	// 由于setName是指针接收的 因此最后也是指针生效 我想表达的是,此时是引用对应着接口
	_, ok = interface{}(&t1).(action)
	fmt.Printf("result: %v\n", ok)
}

type teacher2 struct {
	name string
}

// 签名可以重复
func (p teacher2) talk() {
	fmt.Printf("I am coding %s name\n", p.name)
}

func (p *teacher2) reName() {
	fmt.Printf("I am debuging %s name\n", p.name)
}
