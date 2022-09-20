package main

import (
	"embed"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/jeyrce/gin-app/api"
	"github.com/jeyrce/gin-app/api/middle"
	v1 "github.com/jeyrce/gin-app/api/v1"
	v2 "github.com/jeyrce/gin-app/api/v2"
	"github.com/jeyrce/gin-app/api/view"
	// _ "github.com/jeyrce/gin-app/model"
	c "github.com/jeyrce/gin-app/pkg/conf"
	"github.com/jeyrce/gin-app/pkg/log"
	"github.com/jeyrce/gin-app/task"
)

func init() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
}

//go:embed tmpl
var tmpl embed.FS

//go:embed static
var static embed.FS

//go:embed LICENSE
var license embed.FS

//go:embed favicon.ico
var favicon embed.FS

func main() {
	log.L.Infow("服务开始启动", "listen", viper.GetString(c.MetaListenAddr))
	task.T.Start()
	defer task.T.Stop()
	app := gin.New()
	app.SetHTMLTemplate(template.Must(template.New("tmpl").ParseFS(tmpl, "tmpl/*.html")))
	middle.Registry(app)
	app.HandleMethodNotAllowed = true
	app.NoRoute(api.HTTP404)
	app.NoMethod(api.HTTP405)
	app.StaticFS("/ui", http.FS(static))
	app.StaticFileFS(api.U("/favicon.ico"), "./favicon.ico", http.FS(favicon))
	app.StaticFileFS(api.U("/LICENSE"), "./LICENSE", http.FS(license))
	v1.Registry(app.Group(api.U("/api/v1")))
	v2.Registry(app.Group(api.U("/api/v2")))
	// view 由于需要生成全局路由接口，必须最后加载
	view.Registry(app, app.Group(api.U("/")))
	if err := app.Run(viper.GetString(c.MetaListenAddr)); err != nil {
		log.L.Panic(err)
		os.Exit(1)
	}
}
