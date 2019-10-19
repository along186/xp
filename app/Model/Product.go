package Model

import (
	"time"

	"github.com/pkg/errors"
)

type Product struct {
	Id            int64     `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Image         string    `json:"image"`
	PackageStatus int       `json:"package_status"`
	Status        int       `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

const (
	ProductPackageStatusTrue  int = 1
	ProductPackageStatusFalse int = 2
	ProductStatusNormal       int = 1
	ProductStatusDELETED      int = 2
)

func GetProducts(page int64, limit int64) []Product {
	Products := []Product{}
	db.Model(Product{}).Where("status = ?", ProductStatusNormal).Offset((page - 1) * limit).Limit(limit).Find(&Products)
	return Products
}

func CountProducts() int64 {
	var rows int64 = 0
	db.Model(Product{}).Where("status = ?", ProductStatusNormal).Count(&rows)
	return rows
}

func GetProduct(id int64) (Product, error) {
	product := Product{}
	if db.Model(Product{}).Where("id = ?", id).Where("status = ?", ProductStatusNormal).First(&product).RecordNotFound() == true {
		return product, errors.New("商品不存在")
	}
	return product, nil
}

func SaveProduct(product Product) Product {
	db.Model(Product{}).Create(&product)
	return product
}

func DeleteProduct(id int64) bool {
	product := Product{
		Status:    ProductStatusDELETED,
		UpdatedAt: time.Now(),
	}
	return db.Model(Product{}).Where("id = ?", id).Update(&product).RowsAffected > 0
}
