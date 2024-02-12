package main

import (
	"auth/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	auth_router := router.Group("api/v1/auth/")
	{
		auth_router.POST("signup", controllers.SignUpUser)
	}
	router.Run(":8080")
}
