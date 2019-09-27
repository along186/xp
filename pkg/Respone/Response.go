package Respone

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xp/app/Constant"
)

type Response struct {
	context *gin.Context
}

func Context(c *gin.Context) *Response {
	return &Response{
		context: c,
	}
}

func (r *Response) Success(data interface{}) {
	r.context.JSON(http.StatusOK, gin.H{
		"code" : Constant.SUCCESS,
		"msg" : Constant.GetMsg(Constant.SUCCESS),
		"data" : data,
	})
}

func (r *Response) Error(msg string) {
	r.context.JSON(http.StatusOK, gin.H{
		"code" : Constant.ERROR,
		"msg" : msg,
		"data" : "",
	})
}

func (r *Response) Notice(msg string) {
	r.context.JSON(http.StatusOK, gin.H{
		"code" : Constant.TOAST,
		"msg" : msg,
		"data" : "",
	})
}
