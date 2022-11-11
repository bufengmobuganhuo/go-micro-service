package main

import (
	"github.com/bufengmobuganhuo/go-micro-service/user/handler"
	pb "github.com/bufengmobuganhuo/go-micro-service/user/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user"),
	)

	// Register handler
	pb.RegisterUserHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
