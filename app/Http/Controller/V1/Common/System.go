package Common

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"xp/app/Bill"
	"xp/pkg/Respone"
)

func OrderOpen(c *gin.Context)  {
	open := c.DefaultPostForm("open", "0")
	if com.IsSliceContainsStr([]string{"1","0"}, open) == false {
		Respone.SetContext(c).Notice("参数非法")
	} else {
		if Bill.OrderOpen(open) == true {
			Respone.SetContext(c).Success("")
		} else {
			Respone.SetContext(c).Error("修改失败")
		}
	}
}