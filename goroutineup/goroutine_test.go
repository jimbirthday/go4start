package httpup

import (
	"fmt"
	"testing"
)

//并发指在同一时间内可以执行多个任务
func Test1(T *testing.T) {

	//简单的使用，这里因为使用的test方法，执行后直接结束，所以效果不是很明显，可以切换为main函数执行
	go running("number one")
	//匿名声明
	var a = "number two"
	go func(value string) {
		fmt.Println("i am in goroutine it is go ~~~", value)
	}(a)

}

func running(value string) {
	for {
		fmt.Println("i am in goroutine it is go ~~~", value)
	}
}

func Test2(T *testing.T) {

	chal := make(chan int)

	go func() {
		chal <- 2
		chal <- 1
	}()

	//i := <-chal
	//i1 := <-chal
	//i2, ok := <-chal
	//fmt.Println(i)
	//fmt.Println(i1)
	//if ok {
	//	fmt.Println(i2)
	//} else {
	//	fmt.Println("fatal error: all goroutines are asleep - deadlock!")
	//}

	for data := range chal {
		if data == 0 {
			return
		}
		fmt.Println(data)
	}
}

func Test3(T *testing.T) {
	//10个缓冲的通道
	//缓冲通道被填满时，尝试再次发送数据时发生阻塞
	//缓冲通道为空时，尝试接受数据被阻塞
	ints := make(chan int, 10)

	go printGo(ints)
	for i := 10; i > 0; i-- {
		ints <- i
	}
}

//只能接收的类型
func printGo(c <-chan int) {
	for {
		data := <-c
		fmt.Println(data)
		if data == 0 {
			break
		}
	}
}
