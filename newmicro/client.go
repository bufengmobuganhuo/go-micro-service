package main

import (
	imooc "cap-imooc/proto/cap"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
)

func main() {
	// 实例化
	service := micro.NewService(micro.Name("cap.imooc.client"))
	// 初始化
	service.Init()
	capImooc := imooc.NewCapService("cap.imooc.server", service.Client())
	res, err := capImooc.SayHello(context.TODO(), &imooc.SayRequest{Message: "没有蛀牙"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Answer)
}
