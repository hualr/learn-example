package _2_function

import "testing"

type dog struct {
	name string
}

func setName1(d dog, name string) {
	d.name = name
}

//如果想要改变信息 需要传入的是指针 但凡需要修改 必须走指针
func setName2(d *dog, name string) {
	d.name = name
}

func Test4(t *testing.T) {
	d1 := dog{}

	setName1(d1, "wangcai")
	println("d1的结果为", d1.name)

	setName2(&d1, "wangcai")
	println("d1的结果为", d1.name)
}
