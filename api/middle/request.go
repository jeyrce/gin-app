package middle

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 全局请求记录
func requestLog(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		// Process request
		c.Next()
		// 记录执行时间和一些关键信息
		end := time.Now()
		statusCode := c.Writer.Status()
		logWith := func(msg string, kvs ...interface{}) {}
		switch {
		// 1xx ~ 2xx
		case statusCode < http.StatusMultipleChoices:
			logWith = logger.Infow
		// 3xx
		case statusCode < http.StatusBadRequest:
			logWith = logger.Warnf
		// 4xx - 5xx
		case statusCode >= http.StatusBadRequest:
			logWith = logger.Errorw
		}
		if raw != "" {
			path = path + "?" + raw
		}
		logWith(path, "method", c.Request.Method, "ip", c.ClientIP(),
			"code", statusCode, "duration", fmt.Sprintf("%v", end.Sub(start)),
		)
	}
}
