package repository

import (
	"github/publiusvergilius/mongodb-course/models"
	"slices"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StubUserRepository struct {
	userCollection map[string]*models.User
}

func (u* StubUserRepository) GetAll () ([]*models.User, error) {
	// Here should Be logic to database access
	var users = []*models.User{
		{
			ID: primitive.NewObjectID(),
			Username: "Vini",
			Email: "vini@email.com",
		},
	}

	return users, nil
} 

func TestUserRepository (t *testing.T) {

	want := []*models.User{
		{
			ID: primitive.NewObjectID(),
			Username: "Vini",
			Email: "vini@email.com",
		},
	}

	s := StubUserRepository{}
	got, _ := s.GetAll()	
	
	if !slices.EqualFunc(want, got, compareUserPointers) {
		t.Errorf("want %v, got %v", want, got)
	}

}

func compareUserPointers(p1 *models.User, p2 *models.User) bool {
	if p1 == nil && p2 == nil {
		return true
	}
	if p1 == nil || p2 == nil {
		return false
	}
	return p1.Username == p2.Username && p1.Email == p2.Email
}


