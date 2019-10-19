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

func (m *Session) GetUserInfo(c *gin.Context) (map[string]interface{}, error) {
	token := c.GetHeader("TOKEN")
	url := "http://test5.auth.t.xianghuanji.com/user/getInfo"
	dataUser := make(map[string]string)
	dataUser["token"] = token

	user, _ := Curl.Get(url, dataUser)

	//json str 转map
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(user), &dat); err != nil {
		return dat, err
	}

	userInfo := make(map[string]string)

	userInfo = dat["data"]

	fmt.Println(userInfo["data"])


	//for _, value := range dat {
		//fmt.Println(value)
		//if key == "data" {
		//	userInfo["userid"] = value["id"]
		//}
		//fmt.Printf("%s->%-10s", key, value)
	//}

	return dat, nil
}
