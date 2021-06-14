package client

// 定义全局句柄、需要初始化的变量，然后在initializer子包内单个文件进行初始化

import (
	"github.com/go-redis/redis/v8"
	"github.com/jjonline/go-mod-library/logger"
	"gorm.io/gorm"
)

var (
	Logger *logger.Logger
	Redis  *redis.Client
	DB     *gorm.DB // 内部操作不允许重新赋值，除非清楚您知道现在找干什么！！！
)
