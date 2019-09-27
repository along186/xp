package V1

import (
	"net/http"
	"xp/app/Constant"

	"github.com/gin-gonic/gin"

	"xp/app/Bill"

)

func Index(c *gin.Context)  {

	s := Bill.CheckSystemAvailable()

	if s != true {
		c.JSON(http.StatusOK, gin.H{
			"code" : Constant.TOAST,
			"msg" : Constant.GetMsg(Constant.SYSTEM_UNAVAILABLE),
			"data" : make(map[string]string),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "success",
			"data" : make(map[string]string),
		})
	}
}
