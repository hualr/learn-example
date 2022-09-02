package _1_bean

import (
	"fmt"
	"testing"
)

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

//TODO 完全不理解啊
func Test7(t *testing.T) {
	var dog1 *Dog
	fmt.Println("The first dog is", dog1)
	dog2 := dog1
	if dog2 != nil {
		print("dog2 is not null", dog2)
	}
	// 这里做了一层封装 做的事情是,将类型构建 然后将nil塞到类型中 很难理解!
	var pet1 Pet = dog2
	if pet1 != nil {
		print("pet is not null", pet1)
	}
}
