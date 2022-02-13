package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"

	"github.com/devchrischen/url-shortener/lib/errors"
	"github.com/devchrischen/url-shortener/lib/redis"
)

const (
	DefRateLimiterMaxRetry = 3
	DefRateLimiterPeriod   = 1 * time.Minute
	DefRateLimiterTimes    = 5
)

func RateLimiter(period time.Duration, limit int64) gin.HandlerFunc {

	return func(c *gin.Context) {
		rate := limiter.Rate{
			Period: period,
			Limit:  limit,
		}

		store, err := sredis.NewStoreWithOptions(redis.Client, limiter.StoreOptions{
			Prefix:   redis.APIRateLimiterKeyPrefix,
			MaxRetry: DefRateLimiterMaxRetry,
		})
		if err != nil {
			errors.Error(c, http.StatusInternalServerError, errors.CODE_RATE_LIMITER_ERROR, err)
			return
		}

		middleware := &mgin.Middleware{
			Limiter:        limiter.New(store, rate),
			OnError:        DefaultErrorHandler,
			OnLimitReached: DefaultLimitReachedHandler,
			KeyGetter:      DefaultKeyGetter,
			ExcludedKey:    nil,
		}

		middleware.Handle(c)
	}
}

func DefaultErrorHandler(c *gin.Context, err error) {
	errors.Error(c, http.StatusInternalServerError, errors.CODE_RATE_LIMITER_ERROR, err)
}

func DefaultLimitReachedHandler(c *gin.Context) {
	errors.Throw(
		c,
		errors.NewErr(http.StatusTooManyRequests, errors.CODE_RATE_LIMITER_TOO_MANY),
	)
}

func DefaultKeyGetter(c *gin.Context) string {
	return c.ClientIP()
}
