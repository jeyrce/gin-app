package view

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func viewIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// 全局已注册路由信息，方便调试
func viewPath(e *gin.Engine) gin.HandlerFunc {
	type item struct {
		Path    string
		Method  string
		Handler string
	}
	var (
		paths  = make([]item, 0)
		routes = e.Routes()
	)
	for _, route := range routes {
		paths = append(paths, item{
			Path:    route.Path,
			Method:  route.Method,
			Handler: route.Handler,
		})
	}
	sort.SliceStable(paths, func(i, j int) bool { return paths[i].Path < paths[j].Path })
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "全局路由信息",
			"data":    paths,
		})
	}
}
