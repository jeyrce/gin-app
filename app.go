package main

import (
	"embed"
	"html/template"
	"net/http"
	"os"
	
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	
	"github.com/jeyrce/gin-app/handler"
	"github.com/jeyrce/gin-app/handler/middle"
	v1 "github.com/jeyrce/gin-app/handler/v1"
	v2 "github.com/jeyrce/gin-app/handler/v2"
	"github.com/jeyrce/gin-app/handler/view"
	// _ "github.com/jeyrce/gin-app/model"
	c "github.com/jeyrce/gin-app/pkg/conf"
	"github.com/jeyrce/gin-app/pkg/log"
	"github.com/jeyrce/gin-app/task"
)

func init() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
}

//go:embed lib/tmpl
var tmpl embed.FS

//go:embed lib/static
var static embed.FS

//go:embed LICENSE
var license embed.FS

//go:embed lib/favicon.ico
var favicon embed.FS

func main() {
	log.L.Infow("服务开始启动", "listen", viper.GetString(c.MetaListenAddr))
	task.T.Start()
	defer task.T.Stop()
	app := gin.New()
	app.SetHTMLTemplate(template.Must(template.New("tmpl").ParseFS(tmpl, "lib/tmpl/*.html")))
	middle.Registry(app)
	app.HandleMethodNotAllowed = true
	app.NoRoute(handler.HTTP404)
	app.NoMethod(handler.HTTP405)
	app.StaticFS("/ui", http.FS(static))
	app.StaticFileFS(handler.U("/favicon.ico"), "lib/favicon.ico", http.FS(favicon))
	app.StaticFileFS(handler.U("/LICENSE"), "./LICENSE", http.FS(license))
	v1.Registry(app.Group(handler.U("/api/v1")))
	v2.Registry(app.Group(handler.U("/api/v2")))
	// view 由于需要生成全局路由接口，必须最后加载
	view.Registry(app, app.Group(handler.U("/")))
	if err := app.Run(viper.GetString(c.MetaListenAddr)); err != nil {
		log.L.Panic(err)
		os.Exit(1)
	}
}
