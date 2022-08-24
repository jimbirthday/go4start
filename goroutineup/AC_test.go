package httpup

import (
	"testing"
	"time"
)

func Test2313(t *testing.T) {
	go a()
	time.Sleep(time.Second * 1)
}

func a() {
	defer println("a defer")
	go b()
	println("a done")
}

func b() {
	defer println("b defer")
	panic("b panic")
}
