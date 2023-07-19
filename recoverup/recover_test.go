package recoverup

import (
	"testing"
	"time"
)

func TestRecover(t *testing.T) {
	f1 := func() {
		panic("123")
	}
	f2 := func() {
		panic("123")
	}

	defer func() {
		if err := recover(); err != nil {
			println("321")
		}
	}()
	for {
		select {
		case <-time.After(time.Second * 1):
			func() {
				if time.Second == 1 {

				} else if time.Second == 1 {

				}
				if time.Second == 1 {

				}
				if time.Second == 1 {

				}
			}()
		case <-time.After(time.Second * 2):
			func() {
				if time.Second == 1 {
					println("default")

				} else if time.Second == 1 {

				}
				println("default")

				if time.Second == 1 {

				}
				println("default")

				if time.Second == 1 {
					println("default")

				}
			}()
		default:
			println("default")
		}
	}

}
