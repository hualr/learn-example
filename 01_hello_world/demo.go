package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	/**
	1. 第一个是传入的地址
	2. 第二个是参数
	3. 第三个是默认的值
	4. 第四个是使用说明
	*/
	fmt.Println("init方法执行")
	// 表示的是将传入的参数设置给name
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

// 启动包名必须是main
//go run .\demo.go -name="Rose"
//go run .\demo.go .\hello.go  迭代 放在两个文件得这样
func main() {
	// 解析命令行 不然无效
	flag.Parse()
	hello()
}
