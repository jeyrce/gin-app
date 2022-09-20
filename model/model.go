package model

import (
	"fmt"
	"sync"
	"time"

	c "github.com/jeyrce/gin-app/pkg/conf"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	. "github.com/jeyrce/gin-app/pkg/log"
)

var (
	mysqlTmpl = `%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`
	DB        *gorm.DB // 全局db对象
	once      sync.Once
)

func init() {
	once.Do(func() {
		var driver gorm.Dialector
		if hostPort := viper.GetString(c.DBHostPort); hostPort != "" {
			L.Infof("初始化数据连接: %s", hostPort)
			driver = mysql.Open(fmt.Sprintf(mysqlTmpl, viper.GetString(c.DBUserPass), hostPort, viper.GetString(c.DBDatabase)))
		}
		db, err := gorm.Open(driver, &gorm.Config{
			CreateBatchSize: 100,
			NowFunc: func() time.Time {
				return time.Now().Local()
			},
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
				NoLowerCase:   true,
			},
		})
		if err != nil {
			panic(err)
		}
		DB = db
	})
}

// 类似于 gorm.Model
type meta struct {
	UUID      string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
