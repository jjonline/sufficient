package initializer

import (
	"github.com/jjonline/go-mod-library/logger"
	"github.com/jjonline/golang-backend/conf"
)

func iniLogger() *logger.Logger {
	return logger.New(conf.Config.Log.Level, conf.Config.Log.Path)
}
