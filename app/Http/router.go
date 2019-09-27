package Http

import (
	"github.com/gin-gonic/gin"

	"xp/app/Http/Controller/V1"

	"xp/app/Http/Controller/V1/Order"

	"xp/app/Http/Controller/V1/User"

	"xp/app/Http/Middleware"

)

func InitRouter(e *gin.Engine)  {

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

}