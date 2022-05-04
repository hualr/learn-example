package _1_interface

import "fmt"

type Dog struct {
	name   string
	weight int16
}

// duck理论 如果接收器完美覆盖接口 那么对应的struct就实现了接口
func (dog Dog) eat() string {
	fmt.Println("a Dog eat... ")
	return dog.name
}

func (dog Dog) sleep() string {
	fmt.Println("a Dog sleep ...")
	return dog.name
}

// 因为是指针方法实现的 说明指针结构体才是继承Action 因此最后只能返回指针对象
func (dog *Dog) setName(name string) Action {
	dog.name = name
	return dog
}

func (dog Dog) Name() string {
	return dog.name
}
