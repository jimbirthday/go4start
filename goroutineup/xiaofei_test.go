package httpup

import (
	"testing"
	"time"
)

func product(ch chan<- int, name string) {
	for i := 0; i < 10; i++ {
		ch <- i
		println(name, "product ----", i)
	}
}

func spend(ch <-chan int, name string) {
	for i := 0; i < 10; i++ {
		println(name, "spend ----", <-ch)
	}
}

func Test100(t *testing.T) {
	ints := make(chan int, 10)
	go product(ints, "pro1")
	go product(ints, "pro2")
	go spend(ints, "sp1")
	go spend(ints, "sp2")
	time.Sleep(time.Second * 10)
}
