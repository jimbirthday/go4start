package httpup

import (
	"fmt"
	"net/http"
	"testing"
)

func Test1(T *testing.T) {
	var port = ":8888"
	fmt.Println("hello world")
	handle := http.Handle
	handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("httpup://localhost" + port)
	_ = http.ListenAndServe(port, nil)
}
