package log

import (
	"io"
	"os"
	"sync"

	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	c "github.com/jeyrce/gin-app/pkg/conf"
)

// 全局普通 logger
var (
	config = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder = zapcore.NewJSONEncoder(config)
	writer  io.Writer

	once sync.Once
	L    *zap.SugaredLogger // 全局 logger
)

func init() {
	once.Do(func() {
		// 当设置文件路径时需要进行日志轮转
		switch filename := viper.GetString(c.MetaLogOutput); filename {
		case c.Stdout:
			writer = os.Stdout
		default:
			writer = &lumberjack.Logger{
				Filename:   filename,
				MaxSize:    viper.GetInt(c.MetaLogMaxSize),
				MaxAge:     viper.GetInt(c.MetaLogMaxAge),
				MaxBackups: viper.GetInt(c.MetaLogMaxBackup),
				LocalTime:  true,
				Compress:   viper.GetBool(c.MetaLogCompress),
			}
		}
		logger := zap.New(zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(writer), zap.DebugLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(writer), zap.ErrorLevel),
		), zap.AddCaller())
		defer func() { _ = logger.Sync() }()
		L = logger.Sugar()
		CL = &cronLogger{L}
	})
}
