package main

import (
	"github/publiusvergilius/mongodb-course/controllers"
	"github/publiusvergilius/mongodb-course/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	u := controllers.UsersControlllers{}
	r.GET("/users", u.GetUsers)

	r.Run(":8080")

}
