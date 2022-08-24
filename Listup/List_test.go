package Listup

import (
	"fmt"
	"testing"
)

type List[T any] []T

func (a List[T]) add(val T) {
	fmt.Println("a ", a)
	fmt.Println("a point", &a)
}

func (a *List[T]) add1(val T) *List[T] {
	fmt.Println("a1 ", a)
	i := &a
	fmt.Println("a1 point", &a)
	fmt.Println("a1 value", *a)
	fmt.Println("i value", **i)
	fmt.Println("ii value", *i)
	return a
}

func printLen[T any](list List[T]) {
	fmt.Println(len(list))
}

func TestList(t *testing.T) {
	var ints List[int]
	fmt.Println("var ", ints)
	ints.add(1)
	ints.add1(1)

}
