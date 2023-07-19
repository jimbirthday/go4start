package http

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func Test1(T *testing.T) {
	str1 := "这是一段要被解析的base64"
	encoding := base64.StdEncoding.EncodeToString([]byte(str1))
	fmt.Println(encoding)
	decode, _ := base64.StdEncoding.DecodeString(encoding)
	fmt.Println(string(decode))
}
