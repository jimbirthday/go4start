package http

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
)

//数组是一段固定长度的连续内存区域
func Test1(T *testing.T) {
	var list1 = [...]string{"A", "B", "C"}
	for _, value := range list1 {
		fmt.Println(value)
	}
}

//map使用散列表（hash）实现
func Test2(T *testing.T) {
	var map1 = map[string]int{"a": 1, "b": 2}

	for key, value := range map1 {
		fmt.Println(value)
		fmt.Println(key)
	}
	fmt.Println("------------")

	i, ok := map1["a"]
	if ok {
		fmt.Println(i)
	}
	fmt.Println("------------")

	i1, ok1 := map1["c"]
	if ok1 {
		fmt.Println(i1)
	} else {
		fmt.Println(ok1)
	}
	fmt.Println("------------")
	delete(map1, "b")
	for key, value := range map1 {
		fmt.Println(value)
		fmt.Println(key)
	}
	fmt.Println("------------")

	//并发下的map
	var symcMap1 sync.Map
	symcMap1.Store("a", 1)
	symcMap1.Store("a", 2)

	symcMap1.Range(func(key, value interface{}) bool {
		fmt.Println(key)
		fmt.Println(value)
		return true
	})

	v, loaded := symcMap1.LoadOrStore("a", 3)
	if loaded {
		fmt.Println("the key is alive , the value is ", v)
	}
}

//列表list是一段非连续的存储的容器
func Test3(T *testing.T) {
	list1 := list.New()
	var list2 list.List
	list1.PushBack("a")

	list2.PushBack("b")

	for i := list1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
