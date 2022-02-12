package redis

import (
	"fmt"

	goredis "github.com/go-redis/redis/v8"

	"github.com/devchrischen/url-shortener/config"
)

var (
	Client *goredis.Client
)

func Init() {

	addr := fmt.Sprintf("%v:%v", config.Config.Redis.Host, config.Config.Redis.Port)

	Client = goredis.NewClient(&goredis.Options{
		Addr: addr,
	})
}
