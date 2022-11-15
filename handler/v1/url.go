package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	c "github.com/jeyrce/gin-app/pkg/conf"
)

// 定义该版本swagger信息
func init() {
	SwaggerInfoV1.Version = "1.0"
	SwaggerInfoV1.Host = viper.GetString(c.MetaListenAddr)
	SwaggerInfoV1.InfoInstanceName = "v1"
	SwaggerInfoV1.BasePath = viper.GetString(c.MetaUrlPrefix)
	SwaggerInfoV1.Schemes = []string{"http"}
	SwaggerInfoV1.Title = viper.GetString(c.MetaAppName)
	SwaggerInfoV1.Description = "v1项目文档说明"
}

// 注册该包内部路由信息
func Registry(group *gin.RouterGroup) {
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	group.GET("/version", handleVersion)
}
