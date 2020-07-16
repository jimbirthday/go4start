package http

import (
	"container/list"
	"fmt"
	"testing"
)

//流程控制
func Test1(T *testing.T) {
	var map1 = map[string]int{"a": 1, "b": 2}
	//if 特殊的写法
	if i1, ok1 := map1["c"]; ok1 {
		fmt.Println(i1)
	} else {
		fmt.Println(ok1)
	}
}

//switch
func Test2(T *testing.T) {
	//基本写法
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println("hello")
	case "aaaaa", "bbbb":
		fmt.Println("many case value")
	default:
		fmt.Println("none")
	}
	//特殊写法
	switch {
	case a == "hello":
		fmt.Println("hello")
	default:
		fmt.Println("none")
	}
	//跨越case
	switch a {
	case "hello":
		fmt.Println("hello")
		fallthrough
	case "aaaaa", "bbbb":
		fmt.Println("fallthrough ~~~~~~ ")
	default:
		fmt.Println("none")
	}

}

//goto
func Test3(T *testing.T) {
	list1 := list.New()
	var list2 list.List
	list1.PushBack("a")
	list1.PushBack("b")
	list1.PushBack("c")

	list2.PushBack("b")

	for i := list1.Front(); i != nil; i = i.Next() {
		fmt.Println("this is list 1 - ", i.Value)
		for i1 := list2.Front(); i1 != nil; i1 = i1.Next() {
			if i1.Value == "c" {
				fmt.Println("this is list 2 - ", i1.Value)
				goto GOTOFFF
			}

		}
	}
	//避免继续执行到goto
	return
GOTOFFF:
	fmt.Println("goto here ~~~~`")

}

//使用标签跳出循环
func Test4(T *testing.T) {
	list1 := list.New()
	var list2 list.List
	list1.PushBack("a")
	list1.PushBack("b")
	list1.PushBack("c")

	list2.PushBack("b")
	list2.PushBack("c")
Loop:
	for i := list1.Front(); i != nil; i = i.Next() {
		fmt.Println("this is list 1 - ", i.Value)
		for i1 := list2.Front(); i1 != nil; i1 = i1.Next() {
			if i1.Value == "b" {
				fmt.Println("this is list 2 - ", i1.Value)
				break Loop
			}

		}
	}

}

func Test5(T *testing.T) {
	list1 := list.New()
	var list2 list.List
	list1.PushBack("a")
	list1.PushBack("b")
	list1.PushBack("c")

	list2.PushBack("b")
	list2.PushBack("c")

	for i := list1.Front(); i != nil; i = i.Next() {
		fmt.Println("this is list 1 - ", i.Value)
		for i1 := list2.Front(); i1 != nil; i1 = i1.Next() {
			if i1.Value == "b" {
				fmt.Println("this is list 2 - ", i1.Value)
				continue
			}

		}
	}
Loop:
	for i := list1.Front(); i != nil; i = i.Next() {
		fmt.Println("this is loop list 1 - ", i.Value)
		for i1 := list2.Front(); i1 != nil; i1 = i1.Next() {
			if i1.Value == "b" {
				fmt.Println("this is loop list 2 - ", i1.Value)
				continue Loop
			}

		}
	}

}
