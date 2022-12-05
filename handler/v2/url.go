package v2

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	
	"github.com/jeyrce/gin-app/handler"
	c "github.com/jeyrce/gin-app/pkg/conf"
)

// 定义该版本swagger信息
func init() {
	SwaggerInfoV2.Version = "2.0"
	SwaggerInfoV2.Host = viper.GetString(c.MetaListenAddr)
	SwaggerInfoV2.InfoInstanceName = "v2"
	SwaggerInfoV2.BasePath = handler.U("/api/v2")
	SwaggerInfoV2.Schemes = []string{"http"}
	SwaggerInfoV2.Title = viper.GetString(c.MetaAppName)
	SwaggerInfoV2.Description = "v2项目文档说明"
}

func Registry(group *gin.RouterGroup) {
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	group.GET("/test", handleTest)
}
