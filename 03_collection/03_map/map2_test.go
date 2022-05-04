package main

//1. 修改map不需要传入引用
//2. 注意增删的代码
import (
	"fmt"
	"testing"
)

func Test2(t *testing.T) {
	cityMap := make(map[string]string)

	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//遍历
	printMap(cityMap)

	//删除
	delete(cityMap, "China")

	//修改
	cityMap["USA"] = "DC"

	ChangeValue(cityMap)

	fmt.Println("-------")

	//遍历
	printMap(cityMap)
}

func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key = ", key, "- value = ", value)
	}
}

//cityMap 是一个引用传递
func ChangeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}
