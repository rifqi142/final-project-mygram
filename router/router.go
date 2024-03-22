package router

import "github.com/gin-gonic/gin"

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register",)
	}

	return r
}