package client

// 定义全局句柄、需要初始化的变量，然后在initializer子包内单个文件进行初始化

import (
	"github.com/go-redis/redis/v8"
	"github.com/jjonline/go-mod-library/logger"
)

var (
	Logger *logger.Logger
	Redis  *redis.Client
)
