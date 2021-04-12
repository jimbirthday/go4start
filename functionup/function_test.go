package httpup

import (
	"fmt"
	"testing"
)

/**
go 支持普通函数，匿名函数和闭包
函数本身可以作为值进行传递
支持匿名函数和闭包
函数可以满足接口

所有的传递值都是一块新的内存，无论是传入还是传出都是一块新的内存，成员值都是值传递
表示指针在函数参数值传递中传递的只是指针值，不会复制指针指向的部分
*/

//把方法定义为变量
func Test1(T *testing.T) {
	var i func() = insideFunc
	i()

}

func insideFunc() {
	fmt.Println("this is a inside function")
}

//自定义方法，java函数式编程中的回调函数
func Test2(T *testing.T) {
	list := []string{
		"a",
		"b",
		"c"}

	funcation := []func(string) string{
		stringRemove,
	}

	stringADD(list, funcation)

	for _, value := range list {
		fmt.Println(value)
	}

}

func stringADD(list []string, funcation []func(string) string) {
	for i, _ := range list {
		for _, proc := range funcation {
			list[i] = proc(list[i])
		}
	}

}

func stringRemove(str string) string {
	str = str + "Remove"
	return str
}

//存储方法
func Test3(T *testing.T) {
	map1 := map[string]func(){
		"a": func() {
			fmt.Println("aaaaaa~~~")
		},
	}
	f := map1["a"]
	f()

	f1 := map1["b"]
	f1()
}

type Invoker interface {
	Call(interface{})
}

type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	fmt.Println("from", p)
}

//函数接口实现
/**
此处延展java中多态思想可以定义一个接口多个结构体实现，得到不同的结果
同样的配合上诉go的func特性。这个地方可拥有多样的变种代码
大概可采用了策略模式，可以走向不同的回调函数，回调函数中又可以实现不能的结构体实现函数或者是interface

可以参考http中savehttp代码
此处不多赘述

*/
func Test4(T *testing.T) {
	var i Invoker = new(Struct)
	i.Call("call!!!!!!!")
}

//闭包
/**
闭包指的是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了
自由变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量
*/

func closeBag(value int) func() int {
	return func() int {
		value++
		return value
	}
}

func Test5(T *testing.T) {
	bag := closeBag(1)
	fmt.Println("run 1 bag ~~~", bag())
	fmt.Println("run 2 bag ~~~", bag())
	fmt.Println("run 3 bag ~~~", bag())
	fmt.Println("run 4 bag ~~~", bag())
	f := closeBag(10)
	fmt.Println("run +10 bag ~~~", f())

}

//可变参数函数
func canchangeFunc(str string, a ...interface{}) {
	fmt.Println("this is string    key !!", str)
	for _, value := range a {
		//可以获取类型
		switch value.(type) {
		default:
			fmt.Println("this is interface key !!", value)
		}
	}
}

func Test6(T *testing.T) {
	canchangeFunc("super", "hero", "jim", 4)
}

//延迟语句defer,会按后进先出的顺序执行
func Test7(T *testing.T) {
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
}
