package main

// 1. 这个包名从项目名称开始计算
import "learn-example/09_package/01_init/lib1"

// 2.如果不想导入 只是想实现初始化 需要使用_
import _ "learn-example/09_package/01_init/lib2"

// 初始化的时候一定会被调用
func init() {
	print("这个函数即将开始")
}
func main() {
	lib1.Lib1Test()
}
