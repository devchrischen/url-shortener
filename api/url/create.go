package url

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShortUrl(c *gin.Context) {

	c.String(http.StatusOK, "short url created!")
}
