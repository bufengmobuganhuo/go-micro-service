package main

import (
	"fmt"
	"github.com/bufengmobuganhuo/go-micro-service/product/common"
	"github.com/bufengmobuganhuo/go-micro-service/product/domain/repository"
	service2 "github.com/bufengmobuganhuo/go-micro-service/product/domain/service"
	"github.com/bufengmobuganhuo/go-micro-service/product/handler"
	product "github.com/bufengmobuganhuo/go-micro-service/product/proto/product"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

func main() {
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Fatal(err)
	}

	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.service.product", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8083"),
		micro.Registry(consulRegistry),
		// 绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	mysqlConfig := common.GetMysqlFromConsul(consulConfig, "mysql")
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local", mysqlConfig.User, mysqlConfig.Pwd,
		mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
	db, err := gorm.Open("mysql", connectStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 表名不是负数
	db.SingularTable(false)

	// Initialise service
	service.Init()

	//rp := repository.NewProductRepository(db)
	//rp.InitTable()

	productDataService := service2.NewProductDataService(repository.NewProductRepository(db))

	err = product.RegisterProductHandler(service.Server(), &handler.Product{ProductDataService: productDataService})

	if err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
