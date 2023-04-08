package middlewares

import (
	"net/http"
	"strconv"

	"github.com/ainmtsn1999/orm_jwt_auth/database"
	"github.com/ainmtsn1999/orm_jwt_auth/enums"
	"github.com/ainmtsn1999/orm_jwt_auth/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userData").(jwt.MapClaims)
		userRole := userData["role"].(string)

		if c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			if userRole != enums.Admin {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   "unauthorized",
					"message": "you're not allowed to access this endpoint {userAuth}",
				})
				return
			}
		}

		c.Next()
	}
}

func ProductAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})

			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		userRole := userData["role"].(string)
		product := models.Product{}

		err = db.Select("user_id").First(&product, uint(productId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "not found",
				"message": err.Error(),
			})

			return
		}

		if product.UserID != userID || userRole != enums.Admin {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "you're not allowed to access this endpoint {productAuth}",
			})

			return
		}

		c.Next()
	}
}
