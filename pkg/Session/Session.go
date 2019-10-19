package Session

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"sync"

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

type resData struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data struct{
		Id string `json:"id"`
		UserName string `json:"user_name"`
		Name string	`json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	} `json:"data"`
}
func (m *Session) GetUserInfo(c *gin.Context) map[string]string{
	token := c.GetHeader("TOKEN")
	url := "http://test5.auth.t.xianghuanji.com/user/getInfo"
	dataUser := make(map[string]string)
	dataUser["token"] = token

	user, _ := Curl.Get(url, dataUser)
	//json str 转map
	var dat resData
	json.Unmarshal([]byte(user), &dat)
	userInfo := make(map[string]string)
	userInfo["uid"] = dat.Data.Id
	userInfo["name"] = dat.Data.Name
	userInfo["phone"] = dat.Data.Phone
	userInfo["Email"] = dat.Data.Email
	return  userInfo
}
