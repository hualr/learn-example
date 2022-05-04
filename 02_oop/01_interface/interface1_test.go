package _1_interface

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	myDog := Dog{
		name: "贝贝",
	}
	myDog2 := myDog.setName("欢欢")

	_, ok := interface{}(myDog2).(Action)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)

	// 由于setName是指针接收的 因此最后也是指针生效
	_, ok = interface{}(&myDog2).(Action)
	fmt.Printf("*MyDog implements interface Pet: %v\n", ok)

	myDog2.sleep()
}
