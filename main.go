package main

import (
	"github.com/Divyshekhar/go-jwt-api/controllers"
	"github.com/Divyshekhar/go-jwt-api/intializers"
	"github.com/gin-gonic/gin"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.ConnectToDb()
	intializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Server Running on PORT 3000"})
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	r.Run()
}
