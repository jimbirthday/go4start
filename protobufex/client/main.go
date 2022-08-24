package main

import (
	"context"
	"fmt"
	"go4start/protobufex/pro"
	"go4start/protobufex/token"
	"google.golang.org/grpc"
	"io"
	"time"
)

func main() {
	auth := &token.TokenAuth{
		AppKey:    "123",
		AppSecret: "3211",
	}
	dial, err := grpc.Dial("localhost:8090", grpc.WithInsecure(), grpc.WithPerRPCCredentials(auth))
	if err != nil {
		fmt.Println("Dial err ", err)
		return
	}
	client := pro.NewOrderServiceClient(dial)
	request := &pro.OrderRequest{
		OrderId:   "201907300001",
		TimeStamp: time.Now().Unix(),
	}
	infos, err := client.GetOrderInfos(context.TODO(), request)
	for {
		recv, err := infos.Recv()
		if err == io.EOF {
			continue
		}
		fmt.Println("recv", recv)
	}
}
