package main

import (
	"fmt"
	"github.com/bufengmobuganhuo/go-micro-service/category/common"
	"github.com/bufengmobuganhuo/go-micro-service/category/domain/repository"
	service2 "github.com/bufengmobuganhuo/go-micro-service/category/domain/service"
	"github.com/bufengmobuganhuo/go-micro-service/category/handler"
	category "github.com/bufengmobuganhuo/go-micro-service/category/proto/category"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		// 设置服务地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		// 添加Consul作为注册中心
		micro.Registry(consulRegistry),
	)
	// 获取mysql配置
	mysqlConfig := common.GetMysqlFromConsul(consulConfig, "mysql")

	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local", mysqlConfig.User, mysqlConfig.Pwd,
		mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
	fmt.Println(connectStr)
	// 创建数据库
	db, err := gorm.Open("mysql", connectStr)
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	// 禁止使用复表
	db.SingularTable(false)

	// Initialise service
	service.Init()

	//rp := repository.NewCategoryRepository(db)
	//rp.InitTable()

	categoryDataService := service2.NewCategoryDataService(repository.NewCategoryRepository(db))

	// 注册handler
	err = category.RegisterCategoryHandler(service.Server(), &handler.Category{CategoryDataService: categoryDataService})

	if err != nil {
		log.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
