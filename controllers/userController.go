package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/serenade010/crypto-0819/initializers"
	"github.com/serenade010/crypto-0819/models"
)

func UserCreate(c *gin.Context) {
	//Get data off req body
	var body struct {
		Name     string
		Password string
	}

	c.Bind(&body)
	//Create a user
	user := models.User{Name: body.Name, Password: body.Password}

	// Pass pointer of data to Create
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.Status(400)
		return

	}

	//Return it
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserIndex(c *gin.Context) {
	//Get the user
	var users []models.User
	initializers.DB.Find(&users)

	//Response with them
	c.JSON(200, gin.H{
		"users": users,
	})
}

func UserShow(c *gin.Context) {
	//Get id off URL
	id := c.Param("id")

	//Get the user
	var user models.User
	initializers.DB.First(&user, id)

	//Response with them
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserUpdate(c *gin.Context) {
	//Get the id off the URL
	id := c.Param("id")
	//Get the data off the req body
	var body struct {
		Name     string
		Password string
	}
	c.Bind(&body)
	//Find the user we are updating
	var user models.User
	initializers.DB.First(&user, id)

	//Update it
	initializers.DB.Model(&user).Updates(models.User{Name: body.Name, Password: body.Password})
	//Response with it
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserDelete(c *gin.Context) {
	//Get the id off the URL
	id := c.Param("id")
	//Delete the User
	initializers.DB.Delete(&models.User{}, id)
	//Respond
	c.Status(200)
}
