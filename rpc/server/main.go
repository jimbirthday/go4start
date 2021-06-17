package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const ServiceName = "Hello"

type HelloInterface = interface {
	Hello(s string, s1 *string) error
}

func RegisterService(h HelloInterface) {
	err := rpc.RegisterName(ServiceName, h)
	if err != nil {
		fmt.Println("RegisterName err", err.Error())
		return
	}
}

type HelloService struct {
}

func (c *HelloService) Hello(s string, s1 *string) error {
	*s1 = fmt.Sprintf("hello %s", s)
	println(*s1)
	return nil
}
func main() {
	rpc.RegisterName("Hello", new(HelloService))
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var a io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(a))
	})
	http.ListenAndServe(":1234", nil)
}
