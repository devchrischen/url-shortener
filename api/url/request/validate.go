package request

import (
	"regexp"

	"github.com/devchrischen/url-shortener/lib/errors"
	"github.com/gin-gonic/gin"
)

type RedirectRequest struct {
	HashValue string `uri:"hash_value" binding:"required"`
}

func (r *RedirectRequest) Validate(c *gin.Context) (*RedirectRequest, error) {
	// check binding
	if err := c.ShouldBindUri(r); err != nil {
		return nil, errors.ErrInvalidParams.SetError(err)
	}
	// check param is valid hash
	matched, _ := regexp.MatchString("[A-Za-z0-9]{6}", r.HashValue)
	if !matched {
		return nil, errors.ErrInvalidParams
	}
	return r, nil
}
