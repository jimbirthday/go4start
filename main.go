package main

import "fmt"

type A struct {
	test string
}

func main() {
	a := new(A)
	fmt.Println(a)
}
