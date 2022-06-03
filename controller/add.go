package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Add 上传照片，照片存储到/public/pic中
func Add(c *gin.Context) {
	file, _ := c.FormFile("image")

	saveName := strconv.Itoa(int(time.Now().Unix())) + file.Filename
	Log.Println("存储：", saveName)

	dst := "./public/pic/" + saveName
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)
	f, err := os.OpenFile("./public/pic/pic.xml", os.O_RDWR, 0666)

	if err != nil {
		Log.Println("打开pic.xml文件出错：", err)
	}
	err = deleteLastLine(f)
	if err != nil {
		Log.Println(err)
	}
	f.Close()
	f, err = os.OpenFile("./public/pic/pic.xml", os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_RDWR, 6060)
	var strings = []string{
		"<photo>" + saveName + "</photo>",
		"</pic>",
	}
	addLine(f, strings)
	f.Close()
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
