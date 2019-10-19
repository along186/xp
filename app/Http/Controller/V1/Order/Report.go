package Order

import (
	"time"

	"xp/app/Bill"
	"xp/pkg/Respone"

	"github.com/gin-gonic/gin"
)

func GetEveryBodyOrder(c *gin.Context) {
	date, _ := time.ParseInLocation("2006-01-02", c.DefaultQuery("date", time.Now().Format("2006-01-02")), time.Local)
	data := Bill.GetTodayEveryBodyOrder(date)
	Respone.SetContext(c).Success(data)
}

func ReportCount(c *gin.Context) {
	date, _ := time.ParseInLocation("2006-01-02", c.DefaultQuery("date", time.Now().Format("2006-01-02")), time.Local)
	data := Bill.GetReportCount(date)
	Respone.SetContext(c).Success(data)
}
