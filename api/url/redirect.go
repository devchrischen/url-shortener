package url

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
	// fmt.Printf("req: %+v \n", req)

	// check param is valid hash

	// query db to check if the hash exist

	// query db to check if the hash is not expired

	// return code, message, data(redirect url) as response

	c.String(http.StatusOK, "successfully redirect!")
}
