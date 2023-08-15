package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"jzblog/pkg/setting"
)

// 日志实例化
func NewLogger(setting *setting.AppSettings) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   setting.LogSavePath + setting.LogFileName + setting.LogFileEXT,
		MaxAge:     setting.MaxAge,
		MaxBackups: setting.MaxBackups,
		MaxSize:    setting.LogMaxSize,
		Compress:   setting.Compress,
	}
	write := zapcore.AddSync(&hook)
	//设置日志级别
	var level zapcore.Level
	switch setting.LogLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "jingze",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	//设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		write,
		level,
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段,如：添加一个服务器名称
	filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	logger := zap.New(core, caller, development, filed)
	return logger
}
