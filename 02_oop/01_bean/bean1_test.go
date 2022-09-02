package _1_bean

import (
	"fmt"
	"testing"
)

//01-1 类不需要大写
type dog struct {
	name   string
	weight int
}

// 01-2 构造方法 约定俗称以new开头
func newDog(name string, weight int) dog {
	return dog{name: name, weight: weight}
}

//01 生成对象
func Test1(t *testing.T) {

	// 没有构造器的创建
	dog1 := dog{name: "Dike", weight: 1}
	println(dog1.name, dog1.weight)

	// 直接修改属性 不需要set get这一系列方法.违反闭包原则!
	dog1.name = "Mike"
	println(dog1.name, dog1.weight)

	//这种方法的使用场景是何时呢
	dog2 := new(dog)
	println(dog2.name, dog2.weight)

	// 构造函数 貌似没啥用
	dog3 := newDog("coco", 11)
	fmt.Println(dog3)

	var dog4 dog
	dog4.name = "wanglai"
	println(dog4.name)

}

type animal struct {
	location string
}

type lifeCycle struct {
	state string
}
type ChinaDog struct {
	location string
	dog
	animal
	*lifeCycle
}

// 继承公共方法
func Test2(t *testing.T) {
	dog1 := ChinaDog{location: "beijing"}
	fmt.Printf(dog1.name)

	dog1.weight = 1
	fmt.Println(dog1)

	// 完全形式的继承
	dog2 := ChinaDog{
		location: "HeNan",
		dog:      dog{name: "wangcai", weight: 1},
	}
	println(dog2.location, dog2.name, dog2.weight)

	dog2.animal.location = "1"
	dog2.location = "2"
	println(dog2.animal.location)

	dog3 := ChinaDog{
		location:  "china",
		dog:       dog{name: "小白"},
		animal:    animal{location: "地球"},
		lifeCycle: &lifeCycle{state: "new"},
	}
	println(dog3.location)

	a1 := animal{location: "地球"}
	cycle := lifeCycle{state: "new"}
	dog4 := ChinaDog{
		location:  "china",
		dog:       dog{name: "小白"},
		animal:    a1,
		lifeCycle: &cycle,
	}

	println(dog4.animal.location)
	println(dog4.state)
	// 指针和组合的区别在于 指针会变 和组合不会变 值就是真拷贝,而指针是引用
	a1.location = "月亮"
	cycle.state = "young"
	println(dog4.animal.location)
	println(dog4.state)

}

//TODO
type ChinaDog2 struct {
	location string
	*dog
}
