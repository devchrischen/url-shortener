package url

import (
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

	g.GET("/:hash_value", Redirect)
}
