package controllers

import (
	"net/http"
	"strconv"

	"github.com/ainmtsn1999/orm_jwt_auth/database"
	"github.com/ainmtsn1999/orm_jwt_auth/enums"
	"github.com/ainmtsn1999/orm_jwt_auth/helpers"
	"github.com/ainmtsn1999/orm_jwt_auth/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	product.UserID = userID

	err := db.Debug().Create(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, product)

}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))
	userRole := userData["role"].(string)

	product := models.Product{}

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	if userRole != enums.Admin {
		product.UserID = userID
	}
	
	product.ID = uint(productId)

	err := db.Model(&product).Where("id = ?", uint(productId)).Updates(models.Product{
		Title:       product.Title,
		Description: product.Description,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()

	productId, _ := strconv.Atoi(c.Param("productId"))

	if (db.Delete(&models.Product{}, uint(productId)).RowsAffected == 0) {
		c.JSON(http.StatusNotFound, gin.H{
			"err":     "not found",
			"message": "invalid id or not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

func GetProduct(c *gin.Context) {
	db := database.GetDB()

	productId, _ := strconv.Atoi(c.Param("productId"))

	product := models.Product{}

	err := db.First(&product, productId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, product)
}

func GetAllProduct(c *gin.Context) {
	db := database.GetDB()

	product := []models.Product{}

	err := db.Find(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, product)
}
