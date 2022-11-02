package router

import (
	"gin-test/app/comm"
	"gin-test/app/http/contro/front/post"
	"gin-test/app/http/contro/front/user"
	"github.com/gin-gonic/gin"
)

func Init() {
	v1 := comm.Router.Group("api/v1")
	v1.GET("test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome 1111111gin111111111",
		})
	})

	// 默认的路由
	// user group
	var userGroup = v1.Group("user")
	userGroup.POST("list", user.StartPage)

	// post group
	var postGroup = v1.Group("post")
	postGroup.POST("save", post.Save)
}
