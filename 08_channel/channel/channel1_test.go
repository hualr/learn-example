package channel

import (
	"fmt"
	"testing"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func read(ch chan int) {
	for {
		var b int
		b = <-ch
		fmt.Println(b)
	}
}

func Test1(t *testing.T) {
	intChan := make(chan int, 10)
	go write(intChan)
	go read(intChan)

	time.Sleep(10 * time.Second)
}

func writing(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
		// 当i等于9后就不再打印了，channel被阻塞
		fmt.Printf("channel %d\n", i)
	}
}

//如果不取 就会阻塞
func Test2(t *testing.T) {
	intChan := make(chan int, 10) // 只给10个空间
	writing(intChan)
}

// 检验channel空值
func Test3(t *testing.T) {
	ch := make(chan int, 10) // 只给10个空间
	_, ok := <-ch
	if ok == false {
		fmt.Println("channel空了")
	}
}

// 关闭channel
func Test4(t *testing.T) {
	intChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan <- i
	}

	close(intChan)
}

// case+select取数据
func Test5(t *testing.T) {
	var ch chan int
	ch = make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	// 如果不匹配default 当遇不上就会产生直接是说的悲剧
	/**
	1. select的随机性说明他不是和case一样顺序执行的
	2.
	*/
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		default:
			fmt.Println("get data timeout")
			time.Sleep(time.Second)
		}
	}
}

func Test6(t *testing.T) {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		//close可以关闭一个channel
		close(c)
	}()

	//可以使用range来迭代不断操作channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finished..")
}

// select和go结合使用 否则会报错
func Test7(t *testing.T) {
	c := make(chan string)
	c <- "hello world"
	select {
	/*	case msg := <-c:
		println(msg)*/
	/*	case c <- "happy":
		println(<-c)*/
	default:
		println("channel blocking")
	}
}

//channel select写数据 TODO
func Test8(t *testing.T) {
	c := make(chan string)

	go func() {
		c <- "hello world"
	}()

	go func() {
		c <- "hello world"
	}()
	for {
		select {
		case msg := <-c:
			println(msg)
		case c <- "happy":
			println(<-c)
		default:
			println("channel blocking")
			time.Sleep(time.Second * 10)
		}
	}
}

func Test9(t *testing.T) {
	c := make(chan string)

	go func() {
		c <- "hello world"
	}()

	for {
		select {
		case msg := <-c:
			println(msg)
		case c <- "hello world":
			println(<-c)
		case <-time.After(time.Second * 2):
			println("超时！")

		}
	}
}
