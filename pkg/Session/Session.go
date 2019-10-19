package Session

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"

	"github.com/Unknwon/com"

	"xp/pkg/Curl"
)

type Session struct{}

var s *Session
var once sync.Once

func GetInstance() *Session {
	once.Do(func() {
		s = &Session{}
	})
	return s
}


func (m *Session) GetUserId(c *gin.Context) int {
	userv2 := com.StrTo(c.GetHeader("USERV2")).MustInt()
	return userv2
}

type resData struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data struct{
		Id int `json:"id"`
		UserName string `json:"user_name"`
		Name string	`json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	} `json:"data"`
}
func (m *Session) GetUserInfo(c *gin.Context) map[string]interface{}{
	token := c.GetHeader("TOKEN")
	url := "http://test5.auth.t.xianghuanji.com/user/getInfo"
	dataUser := make(map[string]string)
	dataUser["token"] = token

	user, _ := Curl.Get(url, dataUser)
	fmt.Println(user)
	//json str 转map
	var dat resData
	json.Unmarshal([]byte(user), &dat)
	userInfo := make(map[string]interface{})
	userInfo["uid"] = &dat.Data.Id
	userInfo["name"] = &dat.Data.Name
	userInfo["phone"] = &dat.Data.Phone
	userInfo["Email"] = &dat.Data.Email

	return  userInfo
}
