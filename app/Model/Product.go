package Model

import (
	"time"
)

type Product struct {
	Id            int64     `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Image         string    `json:"image"`
	PackageStatus int       `json:"package_status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func GetProducts(page int64, limit int64) []Product {
	Products := []Product{}
	db.Model(Product{}).Offset((page - 1) * limit).Limit(limit).Find(&Products)
	return Products
}

func CountProducts() int64 {
	var rows int64 = 0
	db.Model(Product{}).Count(&rows)
	return rows
}

func GetProduct(id int64) Product {
	product := Product{}
	db.Model(Product{}).Where("id = ?", id).First(&product)
	return product
}

func SaveProduct(product Product) Product {
	db.Model(Product{}).Create(&product)
	return product
}
