package router

import (
	"github.com/ainmtsn1999/orm_jwt_auth/controllers"
	"github.com/ainmtsn1999/orm_jwt_auth/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication(), middlewares.UserAuth())
		productRouter.GET("/", controllers.GetAllProduct)
		productRouter.GET("/:productId", middlewares.ProductAuth(), controllers.GetProduct)
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productId", controllers.UpdateProduct)
		productRouter.DELETE("/:productId", controllers.UpdateProduct)

	}

	return r
}
