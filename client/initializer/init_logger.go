package initializer

import (
	"github.com/jjonline/go-mod-library/logger"
	"github.com/jjonline/golang-backend/conf"
)

func iniLogger() *logger.Logger {
	// 优先使用命令行参数指定的日志路径和日志级别
	level := conf.Config.Log.Level
	path := conf.Config.Log.Path
	if conf.Cmd.LogLevel != conf.DefaultLogLevel {
		level = conf.Cmd.LogLevel
	}
	if conf.Cmd.LogPath != conf.DefaultLogPath {
		level = conf.Cmd.LogPath
	}

	return logger.New(level, path)
}
