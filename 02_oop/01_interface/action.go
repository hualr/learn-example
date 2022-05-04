package _1_interface

type Action interface {
	eat() string
	sleep() string
	setName(name string) Action
}
