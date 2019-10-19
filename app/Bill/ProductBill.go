package Bill

import (
	"xp/app/Model"
)

func GetProductList(page int64, limit int64) map[string]interface{} {
	data := make(map[string]interface{})
	data["total"] = Model.CountProducts() / limit
	data["page"] = page
	data["limit"] = limit
	data["products"] = Model.GetProducts(page, limit)
	return data
}

func GetProduct(id int64) Model.Product {
	return Model.GetProduct(id)
}

func SaveProduct(product Model.Product) Model.Product {
	return Model.SaveProduct(product)
}
