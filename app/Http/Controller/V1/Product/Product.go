package Product

import (
	"strconv"
	"time"

	"xp/app/Bill"
	"xp/app/Model"
	"xp/pkg/Respone"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	data := Bill.GetProductList(page, limit)
	Respone.SetContext(c).Success(data)
}

func Info(c *gin.Context) {
	if id, err := strconv.ParseInt(c.Param("id"), 10, 64); err == nil {
		data, err := Bill.GetProduct(id)
		if err != nil {
			Respone.SetContext(c).Error(err.Error())
		} else {
			Respone.SetContext(c).Success(data)
		}
	} else {
		Respone.SetContext(c).Error(err.Error())
	}
}

type Form struct {
	Title         string `form:"title" json:"title" binding:"required"`
	Description   string `form:"description" json:"description" binding:"required"`
	Image         string `form:"image" json:"image" binding:"required"`
	PackageStatus int    `form:"package_status" json:"package_status" binding:"required,min=1,max=2"`
}

func Create(c *gin.Context) {
	var product Form
	if err := c.ShouldBind(&product); err != nil {
		Respone.SetContext(c).Error(err.Error())
		return
	}
	p := Model.Product{
		Title:         product.Title,
		Description:   product.Description,
		Image:         product.Image,
		PackageStatus: product.PackageStatus,
		Status:        Model.ProductStatusNormal,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	data := Bill.SaveProduct(p)
	Respone.SetContext(c).Success(data)
}

func Delete(c *gin.Context) {
	if id, err := strconv.ParseInt(c.Param("id"), 10, 64); err == nil {
		data := Bill.DeleteProduct(id)
		Respone.SetContext(c).Success(data)
	} else {
		Respone.SetContext(c).Error(err.Error())
	}
}
