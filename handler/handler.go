package handler

import (
	"net/http"
	"path"
	"strings"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	c "github.com/jeyrce/gin-app/pkg/conf"
)

// 统一url前缀处理
func U(url string) string {
	prefix := strings.TrimSpace(viper.GetString(c.MetaUrlPrefix))
	if prefix == "" {
		prefix = "/"
	}
	return path.Join(prefix, url)
}

func HTTP404(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "notfound",
		"data":    nil,
	})
}

func HTTP405(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "method not allowed",
		"data":    nil,
	})
}

func HTTP403(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "forbidden",
		"data":    nil,
	})
}
