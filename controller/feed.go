package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Feed(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{
		"Title": "zhu",
	})
}
