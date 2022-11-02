package user

import (
	"fmt"
	"gin-test/app/http/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Page struct {
	Page  string `form:"page" json:"page" binding:"required" msg:"页码不能为空"`
	Limit string `form:"limit" json:"limit" binding:"required" msg:"分页数不能为空"`
}

func StartPage(c *gin.Context) {
	//params := c.Query("name")

	//var params1 Page
	//if err := c.ShouldBindJSON(&params1); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": middleware.GetValidMsg(err, params1),
	//	})
	//	return
	//}

	// 参数验证
	params := &Page{}
	fmt.Println(params.Page)
	if err := c.ShouldBind(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": middleware.GetValidMsg(err, params)})
		return
	}

	// 正常输出
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func Save(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
