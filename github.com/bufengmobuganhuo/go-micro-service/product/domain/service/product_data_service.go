package service

import (
	"github.com/bufengmobuganhuo/go-micro-service/product/domain/model"
	"github.com/bufengmobuganhuo/go-micro-service/product/domain/repository"
)

type IProductDataService interface {
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindProductByID(int64) (*model.Product, error)
	FindAllProduct() ([]model.Product, error)
}

// NewProductDataService 创建
func NewProductDataService(productRepository repository.IProductRepository) IProductDataService {
	return &ProductDataService{productRepository}
}

type ProductDataService struct {
	ProductRepository repository.IProductRepository
}

// AddProduct 插入
func (u *ProductDataService) AddProduct(product *model.Product) (int64, error) {
	return u.ProductRepository.CreateProduct(product)
}

// DeleteProduct 删除
func (u *ProductDataService) DeleteProduct(productID int64) error {
	return u.ProductRepository.DeleteProductByID(productID)
}

// UpdateProduct 更新
func (u *ProductDataService) UpdateProduct(product *model.Product) error {
	return u.ProductRepository.UpdateProduct(product)
}

// FindProductByID 查找
func (u *ProductDataService) FindProductByID(productID int64) (*model.Product, error) {
	return u.ProductRepository.FindProductByID(productID)
}

// FindAllProduct 查找
func (u *ProductDataService) FindAllProduct() ([]model.Product, error) {
	return u.ProductRepository.FindAll()
}
