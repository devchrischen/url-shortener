package url

import (
	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {

	g := r.Group("")

	g.POST("/", CreateShortUrl)

	g.GET("/:hash_value", Redirect)
}
