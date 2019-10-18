package Respone

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xp/app/Constant"
)

type Response struct {
	context *gin.Context
}

func SetContext(c *gin.Context) *Response {
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
	return
}

func (r *Response) Error(msg string) {
	r.context.JSON(http.StatusOK, gin.H{
		"code" : Constant.ERROR,
		"msg" : msg,
		"data" : "",
	})
	return
}

func (r *Response) Notice(msg string) {
	r.context.JSON(http.StatusOK, gin.H{
		"code" : Constant.TOAST,
		"msg" : msg,
		"data" : "",
	})
	return
}

func (r *Response) Header(url string)  {
	r.context.Redirect(301, url)
	return
}
