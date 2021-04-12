package httpup

import (
	"fmt"
	"testing"
)

func TestE1(T *testing.T) {
	out := make(chan int)
	in := make(chan int)
	go from(out)
	go squarer(in, out)
	go towhere(in)
}

func from(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v
	}
	close(out)
}
func towhere(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func TestE2(T *testing.T) {
	println(testmirror())
}

func testmirror() string {
	res := make(chan string, 3)
	go func() {
		res <- requestMirro("baidu,com")
	}()
	go func() {
		res <- requestMirro("mirr,com")
	}()
	go func() {
		res <- requestMirro("ttttt,com")
	}()
	return <-res
}
func requestMirro(res string) string {
	return res
}
