package main

import (
	"github.com/serenade010/crypto-0819/initializers"
	"github.com/serenade010/crypto-0819/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
