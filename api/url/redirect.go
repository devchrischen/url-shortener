package url

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	c.String(http.StatusOK, "successfully redirect!")
}
