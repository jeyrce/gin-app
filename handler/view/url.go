package view

import (
	"github.com/gin-gonic/gin"
)

// 公共视图
func Registry(e *gin.Engine, group *gin.RouterGroup) {
	defer e.GET("/path", viewPath(e))
	group.GET("", viewIndex)
}
