package controllers

import (
	"context"
	"github/publiusvergilius/mongodb-course/database"
	"github/publiusvergilius/mongodb-course/models"
	"github/publiusvergilius/mongodb-course/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


var userCollection = database.GetUserCollection

type UsersControlllers struct {
	userService services.UserService
}

func (u *UsersControlllers) GetUsers(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	 _, err := userCollection().Find(ctx, bson.M{})

	// _, err := u.userService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.XML(http.StatusOK, "")
}

func (u *UsersControlllers) CreateUser (c *gin.Context) {
	var user models.User

	if err := c.ShouldBindXML(&user); err != nil {
		c.XML(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := userCollection().InsertOne(ctx, user)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	c.XML(http.StatusCreated, user)
}
