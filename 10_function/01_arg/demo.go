package main

import "fmt"

// 函数式选项模式

type Stu struct {
	Name string
	Age  int
}

// 通常情况下, struct 实例化函数是这样的
// 但是这样的写法不能设置默认值
// 而且字段成员可能发生改变, 如 Age 改为 phoneNum,
func newStu(name string, age int) *Stu {
	return &Stu{
		Name: name,
		Age:  age,
	}
}

// 有了如下这种方式, 称为函数式选项模式

// 1. 定义个一个默认值
var defaultStu = &Stu{
	Name: "Tim",
	Age:  18,
}

// 2. 定义一个函数类型
type OptFunc func(*Stu)

// 3. 利用闭包为每个字段定义一个设置值的函数
func WithName(name string) OptFunc {
	return func(s *Stu) {
		s.Name = name
	}
}

func WithAge(age int) OptFunc {
	return func(s *Stu) {
		s.Age = age
	}
}

func newStu2(opts ...OptFunc) *Stu {
	stu := defaultStu
	for _, o := range opts {
		o(stu)
	}
	return stu
}

func main() {
	x := newStu("yky", 25)
	fmt.Println(x)

	x2 := newStu2()
	fmt.Println(x2)

	x2 = newStu2(
		WithName("yyy"),
	)
	fmt.Println(x2)
}
