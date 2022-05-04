package _3_type

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

func Test1(t *testing.T) {
	// 设置值和设置引用
	dog := Dog{"little pig"}
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)

	// 重要知识点 这里的赋值在java中是做不到的!
	// 这种方式是值传递
	myDog := dog
	dog.SetName("Peter")
	fmt.Printf(myDog.name)
	dog.SetName("monster")

	// 这种方式是引用传递
	myDog2 := &dog
	dog.SetName("Peter")
	fmt.Printf(myDog2.name)

	// 这里因为拿着的是引用 因此 也是会改变!
	var pet Pet = &dog
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())

	dog.SetName("monster")
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())

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
