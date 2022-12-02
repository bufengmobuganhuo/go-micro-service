package handler

import (
	"context"
	"github.com/bufengmobuganhuo/go-micro-service/product/common"
	"github.com/bufengmobuganhuo/go-micro-service/product/domain/model"
	"github.com/bufengmobuganhuo/go-micro-service/product/domain/service"
	product "github.com/bufengmobuganhuo/go-micro-service/product/proto/product"
)

type Product struct {
	ProductDataService service.IProductDataService
}

func (p Product) AddProduct(ctx context.Context, req *product.ProductInfo, resp *product.ResponseProduct) error {
	productAdd := &model.Product{}
	if err := common.SwapTo(req, productAdd); err != nil {
		return err
	}
	productId, err := p.ProductDataService.AddProduct(productAdd)
	if err != nil {
		return err
	}
	resp.ProductId = productId
	return nil
}

func (p Product) FindProductById(ctx context.Context, req *product.RequestId, resp *product.ProductInfo) error {
	prd, err := p.ProductDataService.FindProductByID(req.ProductId)
	if err != nil {
		return err
	}
	if err := common.SwapTo(prd, resp); err != nil {
		return err
	}
	return nil
}

func (p Product) UpdateProduct(ctx context.Context, req *product.ProductInfo, resp *product.Response) error {
	prd := &model.Product{}
	if err := common.SwapTo(req, prd); err != nil {
		return err
	}
	err := p.ProductDataService.UpdateProduct(prd)
	if err != nil {
		return err
	}
	resp.Msg = "更新商品成功"
	return nil
}

func (p Product) DeleteProduct(ctx context.Context, req *product.RequestId, resp *product.Response) error {
	if err := p.ProductDataService.DeleteProduct(req.ProductId); err != nil {
		return err
	}
	resp.Msg = "删除商品成功"
	return nil
}

func (p Product) FindAll(ctx context.Context, all *product.RequestAll, resp *product.AllProduct) error {
	prds, err := p.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}
	for _, prd := range prds {
		prdInfo := &product.ProductInfo{}
		if err := common.SwapTo(prd, prdInfo); err != nil {
			return err
		}
		resp.ProductInfo = append(resp.ProductInfo, prdInfo)
	}
	return nil
}
