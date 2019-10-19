package Http

import (
	"xp/app/Http/Controller/V1/Common"
	"xp/app/Http/Controller/V1/Product"
	"xp/pkg/Respone"

	"github.com/gin-gonic/gin"

	"xp/app/Http/Controller/V1"

	"xp/app/Http/Controller/V1/Order"

	"xp/app/Http/Controller/V1/User"

	"xp/app/Http/Middleware"
)

func InitRouter(e *gin.Engine) {

	// 404
	e.NoRoute(func(context *gin.Context) {
		Respone.SetContext(context).Error("url not found")
	})

	// 首页
	e.GET("/v1/index", V1.Index)

	// 登录
	e.POST("/v1/user/login", User.Login)

	// Oauth授权
	e.POST("/v1/user/auth", User.Auth)

	//
	apiv1 := e.Group("/v1/order")
	apiv1.Use(Middleware.CheckAuthorize())
	{
		// 选餐列表
		apiv1.GET("/list", Order.List)

		// 添加选餐
		apiv1.POST("/add", Order.Add)

		// 选餐详情
		apiv1.GET("/detail/:name", Order.Detail)

		// 选餐详情
		apiv1.POST("/detail", Order.Delete)

	}

	product := e.Group("/v1/products")
	product.Use(Middleware.CheckAuthorize())
	{
		product.GET("", Product.Index)

		product.GET("/:id", Product.Info)

		product.POST("", Product.Create)
	}

	common := e.Group("/common")
	common.Use(Middleware.CheckAuthorize())
	{
		common.POST("/upload", Common.Upload)
	}

}
