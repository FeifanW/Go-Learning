package main

import (
	"Go-Learning/grpc/service"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	user := &service.User{
		Username: "abcd",
		Age:      20,
	}
	// 序列化的过程
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}
	// 反序列化
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser.String())
}
