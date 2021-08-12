package client

// 定义全局句柄、需要初始化的变量，然后在initializer子包内单个文件进行初始化

import (
	"github.com/go-redis/redis/v8"
	"github.com/jjonline/go-lib-backend/guzzle"
	"github.com/jjonline/go-lib-backend/logger"
	"github.com/jjonline/go-lib-backend/memory"
	"gorm.io/gorm"
)

var (
	Logger      *logger.Logger
	Redis       *redis.Client
	DB          *gorm.DB
	MemoryCache *memory.Cache
	Guzzle      *guzzle.Client
)
