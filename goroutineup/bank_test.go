package httpup

import (
	"testing"
)

var Money int

//var mu sync.Mutex

func addMoney(addMoney int) {
	//mu.Lock()
	//defer mu.Unlock()
	println("add ")
	Money = Money + addMoney
	println(Money)
}

func rmMoney(addMoney int) {
	//mu.Lock()
	//defer mu.Unlock()
	println("rm ")
	Money = Money - addMoney
	println(Money)
}

func Test123(t *testing.T) {
	go addMoney(100)
	go rmMoney(50)
	go addMoney(100)
	go rmMoney(50)
}
