package main

/**
len描述的是slice本身的长度
cap表示的是slice当前的容量
cap满了可以自动double扩容
*/
import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func Test3(t *testing.T) {
	numbers := make([]int, 3, 4)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向numbers切片追加一个元素1, numbers len = 4， [0,0,0,1], cap = 5
	numbers = append(numbers, 1)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向一个容量cap已经满的slice 追加元素，自动扩容了
	numbers = append(numbers, 2)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

}

func Test4(t *testing.T) {
	s := []int{1, 2, 3} //len = 3, cap = 3, [1,2,3]

	//1. [0, 2) 注意边界
	s1 := s[0:2] // [1, 2]

	fmt.Println(s1)
	s1[0] = 100

	//2. 这两切片是同源的 也就是说一个改了 另外一个也会改
	fmt.Println("s=", s, "s1=", s1)

	//3. copy 可以将底层数组的slice一起进行拷贝 修改不同源
	s2 := make([]int, 3) //s2 = [0,0,0]

	//将s中的值 依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2)
	s[0] = 1000
	fmt.Println(s2, s)
}

func Test5(t *testing.T) {
	stooges := []string{"Moe", "Larry", "Curly"}
	lang := []string{"php", "golang", "java"}
	stooges = append(stooges, lang...)
}

//...表示的是文字中元素的个数
func Test6(t *testing.T) {
	stooges := [...]string{"Moe", "Larry", "Curly"}
	arr := [...]int{1, 2, 3}
	fmt.Println(len(stooges))
	fmt.Println(len(arr))

}

func Test7(t *testing.T) {
	a := []int{0: 99}
	b := []int{99: 0}
	c := []int{99: 1}
	println(a)
	println(b)
	println(c)
}

func speedTime(handler func() string) {

	t := time.Now()
	handler()
	elapsed := time.Since(t)
	// 利用反射获得函数名
	funcName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	fmt.Println(funcName+"spend time:", elapsed)
}

func Test8(t *testing.T) {
	print([]string{"one", "two", "three"}[rand.Intn(3)])
}
