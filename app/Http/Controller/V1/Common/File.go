package Common

import (
	"crypto/md5"
	"encoding/hex"
	"path"
	"time"

	"xp/pkg/Respone"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		Respone.SetContext(c).Error(err.Error())
		return
	}
	ext := path.Ext(file.Filename)

	ctx := md5.New()
	ctx.Write([]byte(time.Now().String()))
	rand := hex.EncodeToString(ctx.Sum(nil))

	filename := "uploads/" + rand + ext

	if err := c.SaveUploadedFile(file, filename); err != nil {
		Respone.SetContext(c).Error(err.Error())
		return
	}
	data := make(map[string]interface{})
	data["image"] = filename
	Respone.SetContext(c).Success(data)
}
