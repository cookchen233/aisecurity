package middlewares

import (
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"time"
)

func GetMemoryStore(duration time.Duration) *persist.MemoryStore {
	return persist.NewMemoryStore(duration)
}
func CacheMemory(duration time.Duration) gin.HandlerFunc {
	return cache.CacheByRequestURI(GetMemoryStore(duration), duration)
}
