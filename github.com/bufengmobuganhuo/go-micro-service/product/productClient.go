package main

import (
	"context"
	"fmt"
	"github.com/bufengmobuganhuo/go-micro-service/product/common"
	go_micro_service_product "github.com/bufengmobuganhuo/go-micro-service/product/proto/product"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

func main() {
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪
	// 链路追踪
	t, io, err := common.NewTracer("go.micro.service.product.client", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("go.micro.service.product.client"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8084"),
		micro.Registry(consulRegistry),
		// 绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)

	productService := go_micro_service_product.NewProductService("go.micro.service.product", service.Client())

	prd := &go_micro_service_product.ProductInfo{
		ProductName:        "imooc",
		ProductSku:         "cap",
		ProductPrice:       1.1,
		ProductDescription: "imocc-cap",
		ProductCategoryId:  1,
		ProductImage: []*go_micro_service_product.ProductImage{
			{
				ImageName: "cap-image",
				ImageCode: "capimage01",
				ImageUrl:  "xxxxxxx",
			},
		},
		ProductSize: []*go_micro_service_product.ProductSize{
			{
				SizeName: "cap-size",
				SizeCode: "xxxxx",
			},
		},
		ProductSeo: &go_micro_service_product.ProductSeo{
			Id:             0,
			SeoTitle:       "seoTitle",
			SeoKeywords:    "SeoKeywords",
			SeoDescription: "SeoDescription",
			SeoCode:        "SeoCode",
		},
	}

	resp, err := productService.AddProduct(context.TODO(), prd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
