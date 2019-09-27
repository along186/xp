package V1

import (
	"xp/app/Constant"
	"xp/pkg/Respone"

	"github.com/gin-gonic/gin"

	"xp/app/Bill"
)

func Index(c *gin.Context)  {

	s := Bill.CheckSystemAvailable()

	if s != true {
		Respone.Context(c).Notice(Constant.GetMsg(Constant.SYSTEM_UNAVAILABLE))
	} else {
		Respone.Context(c).Success("")
	}
}
