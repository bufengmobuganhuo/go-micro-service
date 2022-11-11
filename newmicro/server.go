package main

import (
	imooc "cap-imooc/proto/cap"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
)

type CapServer struct{}

// 需要实现的服务的逻辑
func (c *CapServer) SayHello(ctx context.Context, req *imooc.SayRequest, resp *imooc.SayResponse) error {
	resp.Answer = "我们的口号是：\"" + req.Message + "\""
	return nil
}

func main() {
	// 创建新的服务
	service := micro.NewService(
		// 服务注册中心通过这个定位服务，所以需要是唯一的
		micro.Name("cap.imooc.server"),
	)
	// 初始化方法
	service.Init()
	// 注册服务
	err := imooc.RegisterCapHandler(service.Server(), new(CapServer))
	if err != nil {
		return
	}
	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
