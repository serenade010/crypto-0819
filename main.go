package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/serenade010/crypto-0819/controllers"
	"github.com/serenade010/crypto-0819/initializers"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	//Testing page
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Success!",
		})
	})

	//api-endpoints
	r.POST("/user", controllers.UserCreate)
	r.GET("/users", controllers.UserIndex)
	r.GET("/user/:id", controllers.UserShow)
	r.PUT("/user/:id", controllers.UserUpdate)
	r.DELETE("user/:id", controllers.UserDelete)

	r.Run()
}
