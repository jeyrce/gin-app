package task

import (
	"time"

	"github.com/robfig/cron/v3"

	"github.com/jeyrce/gin-app/pkg/log"
)

var T = cron.New(cron.WithLocation(time.Local), cron.WithLogger(log.CL))

func init() {
	_, _ = T.AddFunc("0 3 * * *", func() {
		log.L.Infof("测试定时任务")
	})
}
