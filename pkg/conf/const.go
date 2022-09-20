package conf

// 配置名称
const (
	MetaListenAddr   = "meta.listen-addr"
	MetaAppName      = "meta.app-name"
	MetaUrlPrefix    = "meta.url-prefix"
	MetaLogOutput    = "meta.log-output"
	MetaLogMaxSize   = "meta.log-max-size"
	MetaLogMaxAge    = "meta.log-max-age"
	MetaLogMaxBackup = "meta.log-max-backup"
	MetaLogCompress  = "meta.log-compress"
	MetaMediaDir     = "meta.media-dir" // 媒体资源目录

	DBHostPort = "db.host-port"
	DBUserPass = "db.user-pass"
	DBDatabase = "db.database"
)

const (
	Stdout = "/dev/stdout"
)
