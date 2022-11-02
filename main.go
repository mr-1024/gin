package main

import (
	"fmt"
	"gin-test/app/comm"
	"gin-test/app/http/middleware"
	"gin-test/app/orm"
	_ "gin-test/app/orm"
	"gin-test/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	// 路由初始化
	router.Init()
	comm.Router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	// 注册GORM
	gormErr := orm.InitDb()
	if gormErr != nil {
		return
	}

	// 注册异常中间键
	comm.Router.Use(middleware.Recover)
	comm.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome gin",
		})
	})

	// 绑定端口-默认8080
	//err := r.Run(":80")
	err := http.ListenAndServe(":80", comm.Router)
	if err != nil {
		return
	}
}
