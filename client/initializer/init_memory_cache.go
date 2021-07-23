package initializer

import (
	"github.com/jjonline/go-lib-backend/memory"
	"time"
)

func initMemoryCache() *memory.Cache {
	return memory.New(5*time.Minute, 10*time.Minute)
}
