package middleware

import (
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/gin-gonic/gin"

	"github.com/devchrischen/url-shortener/lib/redis"
)

func CacheManager(expire time.Duration) gin.HandlerFunc {
	return cache.CacheByRequestURI(redis.CacheStore, expire)
}
