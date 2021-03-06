package url

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devchrischen/url-shortener/config"
	"github.com/devchrischen/url-shortener/lib/apires"
	"github.com/devchrischen/url-shortener/lib/db"
	"github.com/devchrischen/url-shortener/lib/errors"
	h "github.com/devchrischen/url-shortener/lib/hash"
	surl "github.com/devchrischen/url-shortener/services/url"
)

type CreateRequest struct {
	OriginalUrl string `json:"original_url" form:"original_url" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	// validate request
	var req CreateRequest
	if err := c.ShouldBind(&req); err != nil {
		errors.Throw(c, errors.ErrInvalidParams.SetError(err))
		return
	}

	// check url duplicate => If existing in DB, return the existing short url
	urlService := surl.New(db.DB)
	hashValue, exist, err := urlService.CheckUrlExist(req.OriginalUrl)
	if err != nil {
		errors.Throw(c, err)
		return
	}
	if exist {
		shortUrl := fmt.Sprintf("%s:%s/%s",
			config.Config.BaseURL,
			config.Config.Port,
			hashValue,
		)
		c.JSON(http.StatusConflict, apires.Data{
			Base: apires.Base{
				Code:    errors.CODE_DUPLICATE_KEY,
				Message: "URL already in database!",
			},
			Data: shortUrl,
		})
		return
	}

	// produce unique hash
	for {
		tempHash := h.CreateSixDigitHash()
		hashExist, err := urlService.CheckHashExist(tempHash)
		if err != nil {
			errors.Throw(c, err)
			return
		}
		if !hashExist {
			hashValue = tempHash
			break
		}
	}

	// save original url and hash to database
	if err := urlService.InsertUrlData(hashValue, req.OriginalUrl); err != nil {
		errors.Throw(c, err)
		return
	}

	// return code, message, data(baseUrl + hash) as response
	shortUrl := fmt.Sprintf("%s:%s/%s",
		config.Config.BaseURL,
		config.Config.Port,
		hashValue,
	)
	c.JSON(http.StatusOK, apires.Data{
		Base: apires.Base{
			Code:    errors.CODE_OK,
			Message: "Short URL created!",
		},
		Data: shortUrl,
	})
}
