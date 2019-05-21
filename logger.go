/*
 * @Author: berryberry
 * @LastAuthor: Do not edit
 * @since: 2019-05-10 15:28:45
 * @lastTime: 2019-05-21 22:07:34
 */
package ginfizz

import (
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.SugaredLogger

// logger 模块默认初始化
func initLogger() {
	hook := lumberjack.Logger{
		Filename:   filepath.Join(FizzConfig.App.Log.LogsDirPath, FizzConfig.App.Log.LogRotator.Filename), // 日志文件路径
		MaxSize:    FizzConfig.App.Log.LogRotator.MaxSize,                                                 // megabytes
		MaxBackups: FizzConfig.App.Log.LogRotator.MaxBackups,                                              // 最多保留3个备份
		MaxAge:     FizzConfig.App.Log.LogRotator.MaxAge,                                                  //days
		Compress:   FizzConfig.App.Log.LogRotator.Compress,                                                // 是否压缩 disabled by default
	}

	var level zapcore.Level

	switch strings.ToLower(FizzConfig.App.Log.LogLevel) {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&hook),
		),
		level,
	)

	Logger = zap.New(core).Sugar()
}
