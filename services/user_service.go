package services

import (
	"github/publiusvergilius/mongodb-course/models"
	"github/publiusvergilius/mongodb-course/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (u *UserService) GetAll () ([]*models.User, error) {
	users, _ := u.userRepository.GetAll()

	return users, nil
}
