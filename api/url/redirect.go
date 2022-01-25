package url

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	// validate request

	// check param is valid hash

	// query db to check if the data exist and not expired

	// return code, message, data(redirect url) as response

	c.String(http.StatusOK, "successfully redirect!")
}
