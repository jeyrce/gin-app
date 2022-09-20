package log

import (
	"fmt"

	"go.uber.org/zap"
)

var CL *cronLogger

// 实现 cron 的logger
type cronLogger struct {
	l *zap.SugaredLogger
}

func (c *cronLogger) Info(msg string, keysAndValues ...interface{}) {
	c.l.Infow(msg, keysAndValues...)
}

func (c *cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	c.l.Errorw(fmt.Sprintf("%s: %v", msg, err), keysAndValues...)
}
