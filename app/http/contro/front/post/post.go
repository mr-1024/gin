package post

import (
	"gin-test/app/http/middleware"
	"gin-test/app/orm"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
	Title   string `form:"title" json:"title" binding:"required"  msg:"标题不能不能为空"`
	Content string `form:"content" json:"content" binding:"required"  msg:"内容不能不能为空"`
	Author  string `form:"author" json:"author" binding:"required"  msg:"作者不能不能为空"`
}

func Save(c *gin.Context) {
	request := &Post{}
	if err := c.ShouldBind(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": middleware.GetValidMsg(err, request)})
		return
	}

	orm.DbConn.Create(request)
	// 正常输出
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
