package url

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	OriginalUrl string `json:"original_url" form:"original_url"`
}

func CreateShortUrl(c *gin.Context) {
	// validate request
	var req Request
	if err := c.ShouldBind(&req); err != nil {
		panic(err)
	}
	// check duplicate => Input: url  Output: error

	// produce unique hash => Input: nil  Output: hash string

	// save original url and hash to database => Input: url, hash  Output: error

	// return code, message, data(baseUrl + hash) as response

	c.String(http.StatusOK, "short url created!")
}
