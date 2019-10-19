package Bill

import (
	"math"

	"xp/app/Model"

)

func GetProductList(page int64, limit int64) map[string]interface{} {
	data := make(map[string]interface{})
	data["total"] = math.Ceil(float64(Model.CountProducts()) / float64(limit))
	data["page"] = page
	data["limit"] = limit
	data["products"] = Model.GetProducts(page, limit)
	return data
}

func GetProduct(id int64) (Model.Product, error) {
	return Model.GetProduct(id)
}

func SaveProduct(product Model.Product) Model.Product {
	return Model.SaveProduct(product)
}

func DeleteProduct(id int64) map[string]interface{} {
	res := Model.DeleteProduct(id)
	data := make(map[string]interface{})
	data["delete_success"] = res
	return data
}
