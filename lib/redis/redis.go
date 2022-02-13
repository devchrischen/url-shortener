package redis

import (
	"fmt"

	"github.com/chenyahui/gin-cache/persist"
	goredis "github.com/go-redis/redis/v8"

	"github.com/devchrischen/url-shortener/config"
)

const APIRateLimiterKeyPrefix = "rate-limiter"

var (
	Client     *goredis.Client
	CacheStore *persist.RedisStore
)

func Init() {

	addr := fmt.Sprintf("%v:%v", config.Config.Redis.Host, config.Config.Redis.Port)

	Client = goredis.NewClient(&goredis.Options{
		Addr: addr,
	})

	CacheStore = persist.NewRedisStore(Client)
}
