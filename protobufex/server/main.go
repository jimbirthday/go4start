package main

import (
	"fmt"
	pro2 "go4start/protobufex/pro"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	"time"
)

//订单服务实现
type OrderServiceImpl struct {
}

//获取订单信息s
func (os *OrderServiceImpl) GetOrderInfos(request *pro2.OrderRequest, stream pro2.OrderService_GetOrderInfosServer) error {
	fmt.Println(" 服务端流 RPC 模式")
	md, _ := metadata.FromIncomingContext(stream.Context())
	var appKey string
	var appSecret string

	if key, ok := md["appid"]; ok {
		appKey = key[0]
	}

	if secret, ok := md["appkey"]; ok {
		appSecret = secret[0]
	}

	if appKey != "123" || appSecret != "3211" {
		fmt.Println("Token 不合法")
	}
	orderMap := map[string]pro2.OrderInfo{
		"201907300001": pro2.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": pro2.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": pro2.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}
	for id, v := range orderMap {
		if time.Now().Unix() >= request.TimeStamp {
			fmt.Println("订单序列号ID：", id)
			fmt.Println("订单详情：", v)
			//通过流模式发送给客户端
			stream.Send(&v)
		}
	}
	return nil
}

func main() {
	server := grpc.NewServer()
	pro2.RegisterOrderServiceServer(server, new(OrderServiceImpl))
	listen, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println("Listen err ", err)
		return
	}
	server.Serve(listen)
}
