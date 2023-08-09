package repositories

import "CRUD-Fiber/models"

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id uint) error
}
