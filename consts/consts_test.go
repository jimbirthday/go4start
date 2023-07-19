package consts

import (
	"fmt"
	"testing"
)

const (
	a int8 = 0
	b int8 = 1
	c
)

func TestConst(t *testing.T) {
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
}
