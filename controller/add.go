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
	form, _ := c.MultipartForm()

	files := form.File["image"]
	var saveName []string
	for i, file := range files {
		saveName = append(saveName, strconv.Itoa(int(time.Now().Unix()))+file.Filename)
		Log.Println("存储：", saveName[i])

		dst := "./public/pic/" + saveName[i]
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)
	}

	f, err := os.OpenFile("./public/pic/pic.xml", os.O_RDWR, 0666)

	if err != nil {
		Log.Println("打开pic.xml文件出错：", err)
	}
	err = deleteLastLine(f)
	if err != nil {
		Log.Println(err)
	}
	f.Close()

	f, err = os.OpenFile("./public/pic/pic.xml", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 6060)

	if err != nil {
		Log.Println("打开pic.xml文件失败：", err)
	}

	var strings []string

	for _, name := range saveName {
		strings = append(strings, "<photo>"+name+"</photo>")
	}
	strings = append(strings, "</pic>")
	err = addLine(f, strings)
	if err != nil {
		Log.Println("添加文档失败", err)
	}
	f.Close()

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", saveName))
}
