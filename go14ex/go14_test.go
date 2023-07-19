package go14ex

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test01(t *testing.T) {
	fmt.Println(runtime.GOMAXPROCS(4))
	go f1()
	time.Sleep(time.Second * 1)
	go f2()

	for {
		fmt.Println("main")
		time.Sleep(time.Second)
	}
}

func f1() {
	for {
		//fmt.Println("11111")
		//time.Sleep(time.Second)
	}
}

func f2() {
	for {
		fmt.Println("22222")
		time.Sleep(time.Second)
	}
}
