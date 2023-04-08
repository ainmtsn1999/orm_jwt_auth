package middlewares

import (
	"net/http"
	"github.com/ainmtsn1999/orm_jwt_auth/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthenticated",
				"message": err.Error(),
			})

			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
