package url

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/devchrischen/url-shortener/middleware"
)

func Route(r *gin.RouterGroup) {

	g := r.Group("")

	g.Use(middleware.RateLimiter(
		middleware.DefRateLimiterPeriod,
		middleware.DefRateLimiterTimes,
	))

	g.POST("/", CreateShortUrl)

	g.GET("/:hash_value", middleware.CacheManager(5*time.Minute), Redirect)
}
