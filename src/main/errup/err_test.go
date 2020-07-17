package httpup

import (
	"errors"
	"fmt"
	"testing"
)

//封装一个错误

var myselfErr = errors.New(
	"this is myself err")

func Test1(T *testing.T) {
	if 1 > 0 {
		fmt.Println(myselfErr)
	}
}

type MyErr struct {
	Msg      string
	FileName string
}

//用MyErr实现了error的Error方法
func (e *MyErr) Error() string {
	return fmt.Sprintf(e.FileName, e.Msg)
}

func MyErrorHandler(filename string, msg string) error {
	return &MyErr{msg, filename}
}

func Test2(T *testing.T) {
	if 1 > 0 {
		err := MyErrorHandler("error.go", "this is my error ")
		fmt.Println(err.Error())
	}
}

//使用panic可以主动宕机
func Test3(T *testing.T) {
	recover()
	panic("boom ~~~~~~~~~~~~`")
	fmt.Println("i am alive")
}
