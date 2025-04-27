package repository

import (
	"github/publiusvergilius/mongodb-course/models"
)

/**
 * Parte do Data Access Layer
 * Interface entre os metodos de acesso ao database
 * e os controladores
 * Nao devolve *mongo.Cursor por ser dificil de implementar
 * mas internamente o cursor vai ser injetado ai
*/

type UserRepository interface {
	// Find(ctx context.Context, filter any) ([]models.User, error)
	GetAll() ([]models.User, error)
}

type MongoUserRepository struct {
	// db connection
}

func (m *MongoUserRepository) GetAll () {}
