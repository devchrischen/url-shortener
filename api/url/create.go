package url

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devchrischen/url-shortener/entities/edb"
	"github.com/devchrischen/url-shortener/lib/db"
	h "github.com/devchrischen/url-shortener/lib/hash"
	surl "github.com/devchrischen/url-shortener/services/url"
)

type Request struct {
	OriginalUrl string `json:"original_url" form:"original_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	// validate request
	var req Request
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
	}
	// check duplicate => Input: url  Output: error

	// produce unique hash => Input: nil  Output: hash string
	hashValue := h.CreateSixDigitHash()
	// save original url and hash to database => Input: url, hash  Output: error
	// fmt.Println(db.DB)
	urlService := surl.New(db.DB)
	// fmt.Printf("%+v \n", urlService)
	hash := edb.Hash{
		Value: hashValue,
	}
	// fmt.Println(hash)
	if err := urlService.InsertHash(&hash); err != nil {
		fmt.Println(err)
	}
	url := edb.OriginalUrl{
		HashID: hash.ID,
		Url:    req.OriginalUrl,
	}
	if err := urlService.InsertUrl(&url); err != nil {
		fmt.Println(err)
	}

	// return code, message, data(baseUrl + hash) as response

	c.String(http.StatusOK, "short url created!")
}
