package main

import (
	"fmt"
	"github.com/bufengmobuganhuo/go-micro-service/user/domain/repository"
	service2 "github.com/bufengmobuganhuo/go-micro-service/user/domain/service"
	"github.com/bufengmobuganhuo/go-micro-service/user/handler"
	user "github.com/bufengmobuganhuo/go-micro-service/user/proto/user"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
)

func main() {
	// 创建服务
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	// 服务初始化
	srv.Init()

	// 创建数据库链接
	db, err := gorm.Open("mysql", "root:zy?2526818/micro?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	// 自动创建数据表
	db.SingularTable(true)

	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))

	err = user.RegisterUserHandler(srv.Server(), &handler.User{
		userDataService,
	})
	if err != nil {
		fmt.Println(err)
	}
}
