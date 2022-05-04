package main

import "fmt"

func main() {
	//1. 固定长度的数组 如果没有被初始化 那么就是0
	var myArray1 [10]int
	for i := 0; i < len(myArray1); i++ {
		fmt.Print(myArray1[i], " ")
	}
	fmt.Print("\n===================\n")

	//2. 如果没有补全  那么剩下的也是0
	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}
	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	//3. 数组类型 如果容量不同 还是不同
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	fmt.Println(" ------ ")
	// 4. 可以忽略key或者value
	for _, value := range myArray3 {
		fmt.Println("value = ", value)
	}
}
