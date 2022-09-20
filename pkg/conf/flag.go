package conf

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	pflag.Parse()
	if *showVersion {
		fmt.Println(version())
		os.Exit(0)
	}
	_ = viper.BindPFlags(pflag.CommandLine)
	if *config != "" {
		viper.SetConfigFile(*config)
	}
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

var (
	// meta
	showVersion   = pflag.BoolP("version", "v", false, "查看软件版本")
	config        = pflag.StringP("meta.config-file", "c", "config.yml", "配置文件")
	_             = pflag.String(MetaListenAddr, ":80", "监听地址")
	_             = pflag.String(MetaUrlPrefix, "/", "统一路由前缀")
	_             = pflag.String(MetaAppName, "app", "应用名称")
	metaPlatform  = fmt.Sprintf("%s/%s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())
	metaCommitId  string
	metaVersion   string
	metaBuildAt   string
	metaBranch    string
	metaPoweredBy string
	versionTmpl   = `
powered by https://{{.poweredBy}}
    version:	{{.version}}
    branch:		{{.branch}}
    revision:	{{.commitId}}
    buildAt:	{{.buildAt}}
    platform:	{{.platform}}
`
	MetaVersionMap = map[string]string{
		"version":   metaVersion,
		"commitId":  metaCommitId,
		"branch":    metaBranch,
		"buildAt":   metaBuildAt,
		"platform":  metaPlatform,
		"poweredBy": metaPoweredBy,
	}
	_ = pflag.String(MetaLogOutput, Stdout, "日志输出文件(默认标准输出)")
	_ = pflag.Int(MetaLogMaxAge, 10, "日志最大保留天数")
	_ = pflag.Int(MetaLogMaxSize, 50, "日志单个大小限制(M)")
	_ = pflag.Int(MetaLogMaxBackup, 10, "日志最大轮转数量")
	_ = pflag.Bool(MetaLogCompress, false, "是否启用日志压缩")
	_ = pflag.String(MetaMediaDir, "media/", "媒体文件目录")

	// db
	_ = pflag.String(DBHostPort, "127.0.0.1:3306", "数据库地址:端口")
	_ = pflag.String(DBUserPass, "root:password", "数据库用户:密码")
	_ = pflag.String(DBDatabase, "dbname", "数据库名")
)

// 输出应用版本信息
func version() string {
	t := template.Must(template.New("version").Parse(versionTmpl))

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "version", MetaVersionMap); err != nil {
		panic(err)
	}
	return strings.TrimSpace(buf.String())
}
