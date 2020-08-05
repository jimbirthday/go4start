package stringup

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

//变量
var (
	a string
	b int
	c []int
	d func() bool // 声明一个返回值为布尔类型的函数变量，这种形式一般用于回调函数，即将函数以变量的形式保存下来，在需要的时候重新调用这个函数
	e struct {
		x int
	}
)

//常量
const (
	//这里会进行位移运算，分别得到1,2，4，8的值
	AB = 1 << iota
	BC
	CD
	DE
)

func Test1(T *testing.T) {
	var str1 = "超级英雄吉姆"
	var str2 = "superherojim"

	//此处输出的是ASCII编码的长度
	fmt.Println(len(str1))
	fmt.Println(len(str2))

	//需要进行编码来获取utf-8的长度
	fmt.Println(utf8.RuneCountInString(str1))
	fmt.Println(utf8.RuneCountInString(str2))

	//截取字符串
	comma := strings.Index(str1, "英雄")
	fmt.Println(comma, str1[:comma])

	//修改字符串，在go中字符串是不可变的，只能重新copy一个修改
	b1 := []byte(str2)
	for i := 5; i < len(str2); i++ {
		b1[i] = 's'
	}
	fmt.Println(string(b1))

	//连接字符串
	var bytesBuffter bytes.Buffer
	bytesBuffter.WriteString(str1)
	bytesBuffter.WriteString(str2)
	fmt.Println(bytesBuffter.String())

}

func Test2(T *testing.T) {
	fmt.Println(AB)
	fmt.Println(BC)
	fmt.Println(CD)
	fmt.Println(DE)
}

func Test3(T *testing.T) {
	i2 := []byte("this is jim")
	i := make([]byte, len(i2))
	strncpy(i, i2, len(i))
	println(string(i))
}

func strncpy(dist []byte, src []byte, size_t int) {
	for i := 0; i < size_t; i++ {
		dist[i] = src[i]
	}
}

func add(z int) {
	println("z ----", z)
	z += 200
}
func Derfer(x, y int) (z int) {
	defer func() {
		println("derfer", z)
		z += 100
	}()
	z = x + y
	add(z)
	return z
}
func Test4(T *testing.T) {

	derfer1 := Derfer(1, 1)
	println("main  ---", derfer1)
}
func Test5(T *testing.T) {
	i := []byte("this is jim")
	i2 := make([]byte, len(i))
	strncpy1(i2, i, len(i))
	println(string(i2))
}
func strncpy1(l1 []byte, l2 []byte, size int) {
	for i := 0; i < size; i++ {
		l1[i] = l2[i]
	}

}

func Test6(T *testing.T) {
	var i = 1
	j := &i
	println(j)
	*j = 2
	println(i)

}
