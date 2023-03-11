package main

import (
	"Go-Learning/grpc/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// 无认证grpc http2 https
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("服务端出错，连接不上", err)
	}
	defer conn.Close()
	prodClient := service.NewProdServiceClient(conn)

	request := &service.ProductRequest{
		ProdId: 123,
	}

	stockRespons, err := prodClient.GetProductStock(context.Background(), request)
	if err != nil {
		log.Fatal("查询库存出错", err)
	}
	fmt.Println("查询成功", stockRespons)
}
