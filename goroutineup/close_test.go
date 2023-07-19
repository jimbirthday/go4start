package httpup

import (
	"testing"
	"time"
)

func Test23(t *testing.T) {
	f := make(chan int)
	go func() {
		f <- 1
		close(f)
	}()

	println(<-f)
	println(<-f)
	f <- 2
}
func Test33(t *testing.T) {
	f := make(chan int)
	go func() {
		f <- 1
		f <- 2
		close(f)
	}()

	println(<-f)
	println(<-f)
	println(<-f)
	f <- 2
}

//这里要注意close的位置
func Test34(t *testing.T) {
	f := make(chan int, 2)
	go func() {
		f <- 1
		f <- 2
		close(f)
	}()
	time.Sleep(time.Second * 2)
	f <- 2
	println(<-f)
	println(<-f)
	println(<-f)

}
func Test35(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer println(i)
	}

}
