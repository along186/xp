package bootstrap

import (
	"github.com/gin-gonic/gin"
	"xp/pkg/Setting"

	"xp/app/Http"

)

func Run() *gin.Engine {

	// 1.初始化gin框架引擎
	r := gin.Default()

	// 2.设置运行模式
	gin.SetMode(Setting.RunMode)

	// 3.加载路由
	Http.InitRouter(r)

	// 4.返回gin.Engine
	return r
}