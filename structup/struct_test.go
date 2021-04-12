package httpup

import (
	"fmt"
	"testing"
)

//结构体可以被实例化，从New或者&构造的结构体的类型是结构体的指针
//在加上*以后，
//1.传递的其实是参数的指针
//2.内部变量传递的是变量指针指向的值，所以指针的地址永不改变，改变的只是指针指向的值
//一个传的是指针，一个传的是指针值
type abc struct {
	A int
}

func Test1(T *testing.T) {
	a := new(abc)
	fmt.Println(&a)
	fmt.Println(*a)
	fmt.Println(a)

	a.set1(1)
	fmt.Println(&a)
	fmt.Println(*a)
	fmt.Println(a)

}

func (a *abc) set1(value int) {
	fmt.Println("````````````")
	fmt.Println(&a)
	fmt.Println(*a)
	fmt.Println(a)

	a.A = value
	fmt.Println(&a)
	fmt.Println(*a)
	fmt.Println(a)
	fmt.Println("````````````")

}

func (a abc) set(value int) {
	fmt.Println("````````````")
	fmt.Println(&a)
	fmt.Println(a)
	a2 := &a
	fmt.Println(&a2)

	a.A = value
	fmt.Println(&a)
	fmt.Println(a)
	fmt.Println(a2)
	fmt.Println("````````````")

}
func Test2(T *testing.T) {
	a := new(abc)
	//fmt.Println(&a)
	//fmt.Println(*a)
	//fmt.Println(a)

	a.set(1)
	//fmt.Println(&a)
	//fmt.Println(*a)
	fmt.Println(a)

}

func NewABC(v int) *abc {
	return &abc{1}
}

func Test3(T *testing.T) {
	a := NewABC(1)
	fmt.Println(&a)
	fmt.Println(*a)
	fmt.Println(a)
	fmt.Println(a.A)

}

func NewABC1(v int) abc {
	a := abc{1}
	fmt.Println(&a)
	return a
}

func Test4(T *testing.T) {
	a := NewABC1(1)
	fmt.Println(&a)
	fmt.Println(a)
	fmt.Println(a.A)

}

/**
可以利用特性来构造一个事件系统
相同的函数名不管是结构体的函数还是自定义函数
都会被同时调用
*/
