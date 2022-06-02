package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//上传照片，照片存储到/public/pic中
func Add(c *gin.Context) {
	file, _ := c.FormFile("file")
	Log.Println("存储：", file.Filename)

	dst := "./public/pic" + file.Filename
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
