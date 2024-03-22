package router

import (
	"final-project-mygram/controller"
	"final-project-mygram/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.Register)
		userRouter.POST("/login", controller.Login)
		userRouter.PUT("/", middleware.Authentication(), controller.UpdateDataUser)
		userRouter.DELETE("/", middleware.Authentication(), controller.DeleteUser)
	}

	return r
}