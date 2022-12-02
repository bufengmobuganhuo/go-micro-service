package model

type Product struct {
	ID                 int64          `gorm:"primary_key;not_null;auto_increment"`
	ProductName        string         `json:"product_name"`
	ProductSku         string         `gorm:"unique_index;not_null" json:"product_sku"`
	ProductPrice       float64        `json:"product_price"`
	ProductDescription string         `json:"product_description"`
	ProductImage       []ProductImage `gorm:"foreignKey:ImageProductID" json:"product_image"`
	ProductSize        []ProductSize  `gorm:"foreignKey:SeoProductID" json:"product_size"`
	ProductSeo         ProductSeo     `gorm:"foreignKey:SizeProductID" json:"product_seo"`
}
