package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/serenade010/crypto-0819/initializers"
	"github.com/serenade010/crypto-0819/models"
)

func MlModelCreate(c *gin.Context) {
	//Get data off req body
	var body struct {
		Name           string
		Ratio_of_train float32
		Look_back      int
		Learning_rate  float32
		Epochs         int
		Batch_size     int
		UserID         uint
	}

	c.Bind(&body)
	//Create a ML Model
	model := models.MlModel{Name: body.Name, Ratio_of_train: body.Ratio_of_train, Look_back: body.Look_back, Learning_rate: body.Learning_rate, Epochs: body.Epochs, Batch_size: body.Batch_size, UserID: body.UserID}

	// Pass pointer of data to Create
	result := initializers.DB.Create(&model)
	if result.Error != nil {
		c.Status(400)
		return

	}

	//Return it
	c.JSON(200, gin.H{
		"model": model,
	})
}

func MlModelIndex(c *gin.Context) {
	//Get the user
	var models []models.MlModel
	initializers.DB.Find(&models)

	//Response with them
	c.JSON(200, gin.H{
		"models": models,
	})
}

func MlModelShow(c *gin.Context) {
	//Get id off URL
	id := c.Param("id")

	//Get the user
	var model models.MlModel
	initializers.DB.First(&model, id)

	//Response with them
	c.JSON(200, gin.H{
		"user": model,
	})
}

func MlModelUpdate(c *gin.Context) {
	//Get the id off the URL
	id := c.Param("id")
	//Get the data off the req body
	var body struct {
		Name           string
		Ratio_of_train float32
		Look_back      int
		Learning_rate  float32
		Epochs         int
		Batch_size     int
		UserID         uint
	}
	c.Bind(&body)
	//Find the user we are updating
	var model models.MlModel
	initializers.DB.First(&model, id)

	//Update it
	initializers.DB.Model(&model).Updates(models.MlModel{Name: body.Name, Ratio_of_train: body.Ratio_of_train, Look_back: body.Look_back, Learning_rate: body.Learning_rate, Epochs: body.Epochs, Batch_size: body.Batch_size, UserID: body.UserID})
	//Response with it
	c.JSON(200, gin.H{
		"user": model,
	})
}

func MlModelDelete(c *gin.Context) {
	//Get the id off the URL
	id := c.Param("id")
	//Delete the User
	initializers.DB.Delete(&models.MlModel{}, id)
	//Respond
	c.Status(200)
}
