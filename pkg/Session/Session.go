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

func (m *Session) GetUserInfo(c *gin.Context) map[string]string {
	token := c.GetHeader("TOKEN")
	url := "http://test5.auth.t.xianghuanji.com/user/getInfo"
	dataUser := make(map[string]string)
	dataUser["token"] = token

	user, _ := Curl.Get(url, dataUser)

	//json str 转map
	var dat map[string]string
	if err := json.Unmarshal([]byte(user), &dat); err != nil {
		return dat
	}
	fmt.Println("============")

	userInfo := make(map[string]string)

	userInfo["data"] = dat["data"]

	fmt.Println(userInfo["data"])

	return userInfo
}
