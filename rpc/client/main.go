package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	tcp, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("DialTCP err", err.Error())
		return
	}
	codec := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(tcp))
	res := ""
	err = codec.Call("Hello.Hello", "world", &res)
	if err != nil {
		fmt.Println("Call err", err.Error())
		return
	}
	fmt.Println("res : ", res)
}
