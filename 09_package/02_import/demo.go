package main

// 1. 取得别名
import f "fmt"

// 不再调用包名
import . "learn-example/09_package/02_import/lib1"

// 2. 完整获取
func main() {
	f.Println("haha")
	Test()
}
