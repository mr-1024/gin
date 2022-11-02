package comm

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
)

// Router 初始化路由
var Router = gin.Default()

//var Router = mux.NewRouter()
