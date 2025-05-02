package repository

import (
	"github/publiusvergilius/mongodb-course/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StubUserRepository struct {
	userCollection map[string]*models.User
}

func (u StubUserRepository) GetAll () ([]*models.User, error) {
	// Here should Be logic to database access
	var users = []*models.User{
		{
			ID: primitive.NewObjectID(),
			Username: "Vini",
			Email: "vini@email.com",
		},
		{
			ID: primitive.NewObjectID(),
			Username: "Noe",
			Email: "noe@email.com",
		},
	}

	return users, nil
} 
