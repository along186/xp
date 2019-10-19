package V1

import (
	"fmt"
	"github.com/Unknwon/com"
	"xp/pkg/Respone"
	"xp/pkg/Session"

	"github.com/gin-gonic/gin"

	"xp/app/Bill"
)

func Index(c *gin.Context)  {

	returnData := make(map[string]int)
	returnData["orderCheck"] = 0
	returnData["systemCheck"] = 0

	// 1.检查用户是否登录
	userId := com.StrTo(Session.GetInstance().GetUserUid(c)).MustInt()
	if userId != 0 { // 已登录
		orderCheck := Bill.CheckTodaykHasOrdered(userId)
		if orderCheck == true { // 已点餐
			returnData["orderCheck"] = 1
		}
	}

	// 2.检查系统是否可用
	systemCheck := Bill.CheckSystemAvailable()
	if systemCheck != true {
		returnData["systemCheck"] = 0
	} else {
		returnData["systemCheck"] = 1
	}

	Respone.SetContext(c).Success(returnData)

}
