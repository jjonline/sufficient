package initializer

import (
	"github.com/jjonline/go-mod-library/logger"
	"github.com/jjonline/golang-backend/config"
)

func iniLogger() *logger.Logger {
	return logger.New(config.Config.Log.Level, config.Config.Log.Path)
}
