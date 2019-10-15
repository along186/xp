package Session

import (
	"github.com/gin-gonic/gin"
	"sync"

	"github.com/Unknwon/com"

)

type Session struct {}

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
