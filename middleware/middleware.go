package middleware

import (
	"final-project-mygram/helper"
	"final-project-mygram/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func (c *gin.Context) {
		verifyToken, err := helper.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.CreateResponse(false, nil, err.Error(), "Unauthenticated"))
			return
		}
		c.Set("userInfo", verifyToken)
		c.Next()
	}
}