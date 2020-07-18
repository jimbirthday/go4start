package httpup

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

//模拟RPC客户端的请求和接受消息
func RPCClient(ch chan string, req string) (string, error) {
	//模拟请求，发送的客户端
	ch <- req
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("Time out")
	}
}

//RPC服务端接受客户端请求和回应
func RPCServer(ch chan string) {
	for {
		data := <-ch
		fmt.Println(data)
		//模拟超时
		time.Sleep(time.Second * 2)
		ch <- "我收到了请求"
	}
}

func Test4(T *testing.T) {
	ch := make(chan string)
	//启动监听
	go RPCServer(ch)
	//请求服务器
	client, err := RPCClient(ch, "我发起了请求")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(client)
	}

}
