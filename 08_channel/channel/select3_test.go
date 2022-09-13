package channel

import "fmt"

// 尽量返回两个 用于发现一行信息
func SelectAssign(c chan string) {
	select {
	case <-c: //0个变量
		fmt.Println("0")
	case d := <-c: //1个变量
		fmt.Printf("1: received %s \n", d)
	case d, ok := <-c: //2个变量
		if !ok {
			fmt.Printf("no data found")
			break
		}
		fmt.Printf("2: received %s \n", d)
	}
}
