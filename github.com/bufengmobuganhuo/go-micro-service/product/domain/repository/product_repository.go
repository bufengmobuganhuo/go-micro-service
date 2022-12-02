package repository

import (
	"github.com/bufengmobuganhuo/go-micro-service/product/domain/model"
	"github.com/jinzhu/gorm"
)

type IProductRepository interface {
	InitTable() error
	FindProductByID(int64) (*model.Product, error)
	CreateProduct(*model.Product) (int64, error)
	DeleteProductByID(int64) error
	UpdateProduct(*model.Product) error
	FindAll() ([]model.Product, error)
}

// NewProductRepository 创建productRepository
func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDb: db}
}

type ProductRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *ProductRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Product{}, &model.ProductImage{}, &model.ProductSize{}, &model.ProductSeo{}).Error
}

// FindProductByID 根据ID查找Product信息
func (u *ProductRepository) FindProductByID(productID int64) (product *model.Product, err error) {
	product = &model.Product{}
	return product, u.mysqlDb.Preload("ProductImage").Preload("ProductSize").
		Preload("ProductSeo").Find(product, productID).Error
}

//创建Product信息
func (u *ProductRepository) CreateProduct(product *model.Product) (int64, error) {
	return product.ID, u.mysqlDb.Create(product).Error
}

//根据ID删除Product信息
func (u *ProductRepository) DeleteProductByID(productID int64) error {
	// 开启事务，删除多个表
	tx := u.mysqlDb.Begin()
	// 如果发生问题则需要回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Unscoped().Where("image_product_id = ?", productID).Delete(&model.ProductImage{},
		productID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("size_product_id = ?", productID).Delete(&model.ProductSize{},
		productID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("seo_product_id = ?", productID).Delete(&model.ProductSeo{},
		productID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("id = ?", productID).Delete(&model.Product{}, productID).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// UpdateProduct 更新Product信息
func (u *ProductRepository) UpdateProduct(product *model.Product) error {
	return u.mysqlDb.Model(product).Update(product).Error
}

// FindAll 获取结果集
func (u *ProductRepository) FindAll() (productAll []model.Product, err error) {
	return productAll, u.mysqlDb.Preload("ProductImage").Preload("ProductSize").
		Preload("ProductSeo").Find(&productAll).Error
}
