// services/user_service_impl.go
package services

import (
	"CRUD-Fiber/models"
	"CRUD-Fiber/repositories"
)

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserServiceImpl(userRepository repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository: userRepository}
}

func (s *UserServiceImpl) GetAllUsers() ([]models.User, error) {
	return s.userRepository.GetAll()
}

func (s *UserServiceImpl) GetUserByID(id uint) (*models.User, error) {
	return s.userRepository.GetByID(id)
}

func (s *UserServiceImpl) CreateUser(user *models.User) error {
	return s.userRepository.Create(user)
}

func (s *UserServiceImpl) UpdateUser(user *models.User) error {
	return s.userRepository.Update(user)
}

func (s *UserServiceImpl) DeleteUser(id uint) error {
	return s.userRepository.Delete(id)
}
