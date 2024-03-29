package Order

import (
	"fmt"
	"math/rand"
	"time"

	"xp/app/Bill"
	"xp/app/Model"
	"xp/pkg/Respone"
	"xp/pkg/Session"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	//"xp/pkg/Session"
)

func List(c *gin.Context) {

}

func Detail(c *gin.Context) {

}

type Form struct {
	ProductId int64 `form:"product_id" json:"product_id" binding:"required"`
}

func Add(c *gin.Context) {
	var product Form
	if err := c.ShouldBind(&product); err != nil {
		Respone.SetContext(c).Error(err.Error())
		return
	}

	//userId := Session.GetInstance().GetUserId(c)
	userInfo := Session.GetInstance().GetUserInfo(c)
	userId := com.StrTo(userInfo["uid"]).MustInt()
	UserName := userInfo["name"]

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	orderNo := time.Now().Format("20060102150405") + com.StrTo(vcode).String()

	p := Model.Order{
		OrderNo:   orderNo,
		ProductId: product.ProductId,
		Uid:       userId,
		UserName:  UserName,
		Status:    Model.ORDER,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	data, err := Bill.SaveOrder(p)
	if err != nil {
		Respone.SetContext(c).Error(err.Error())
	} else {
		Respone.SetContext(c).Success(data)
	}
}

func Delete(c *gin.Context) {
	userInfo := Session.GetInstance().GetUserInfo(c)
	userId := com.StrTo(userInfo["uid"]).MustInt()
	data := make(map[string]interface{})
	data["delete_success"] = Bill.DeleteOrderByUid(userId)
	Respone.SetContext(c).Success(data)
}
