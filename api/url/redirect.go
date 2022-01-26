package url

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"

	"github.com/devchrischen/url-shortener/entities/edb"
	"github.com/devchrischen/url-shortener/lib/db"
	t "github.com/devchrischen/url-shortener/lib/time"
	surl "github.com/devchrischen/url-shortener/services/url"
)

type redirectRequest struct {
	HashValue string `uri:"hashValue" binding:"required"`
}

func Redirect(c *gin.Context) {
	// validate request
	var req redirectRequest
	if err := c.ShouldBindUri(&req); err != nil {
		fmt.Println(err)
	}

	// check param is valid hash
	match, _ := regexp.MatchString("[A-Za-z0-9]{6}", req.HashValue)
	if !match {
		fmt.Println("Invalid hash value!")
	}
	// query db to check if the hash exist
	urlService := surl.New(db.DB)
	hash := edb.Hash{}
	if err := urlService.GetHashRecord(&hash, req.HashValue); err != nil {
		fmt.Println(err)
	}
	// check if the hash is not expired
	expired := t.CheckHashExpired(hash.CreatedAt)
	if expired {
		fmt.Println("The short url was expired!")
	}
	// return code, message, data(redirect url) as response

	c.String(http.StatusOK, "successfully redirect!")
}
