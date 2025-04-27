package services

import (
	"github/publiusvergilius/mongodb-course/repository"
	"testing"
)

func TestUserService (t *testing.T) {
	userService := UserService{
		userRepository: repository.StubUserRepository{},
	}

	stubUsers, err := userService.GetAll()
	if err != nil {
		t.Error("was not told to err: ", err.Error())
	}

	want := 2
	got := len(stubUsers)

	if want != got {
		t.Errorf("want %d users, got %d", want, got)
	}

}
