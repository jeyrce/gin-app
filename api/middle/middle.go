package middle

import (
	"github.com/gin-gonic/gin"
	"github.com/jeyrce/gin-app/pkg/log"
)

var recovery = gin.Recovery()

// 注册所有中间件
func Registry(e *gin.Engine) {
	e.Use(
		recovery,          // 异常崩溃恢复
		requestLog(log.L), // 全局请求日志
	)
}
