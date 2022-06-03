package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"simple-server/controller"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")
	r.LoadHTMLGlob("./public/page/*")

	Router := r.Group("/")

	Router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", gin.H{
			"title": "Main website",
		})
	})
	Router.GET("/upload/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"title": "图片上传页面",
		})
	})

	Router.GET("/show/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "show.html", gin.H{
			"title": "所有图片显示页面",
		})
	})

	Router.POST("/add/", controller.Add)

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	//apiRouter.GET("/user/", controller.UserInfo)
	//apiRouter.POST("/user/register/", controller.Register)
	//apiRouter.POST("/user/login/", controller.Login)
	//apiRouter.POST("/publish/action/", controller.Publish)
	//apiRouter.GET("/publish/list/", controller.PublishList)

	//// extra apis - I
	//apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	//apiRouter.GET("/favorite/list/", controller.FavoriteList)
	//apiRouter.POST("/comment/action/", controller.CommentAction)
	//apiRouter.GET("/comment/list/", controller.CommentList)

	//// extra apis - II
	//apiRouter.POST("/relation/action/", controller.RelationAction)
	//apiRouter.GET("/relation/follow/list/", controller.FollowList)
	//apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	file, err := os.OpenFile("./log/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		controller.Log.Out = file
	} else {
		controller.Log.Info("Failed to log to file, using default stderr")
	}

}
