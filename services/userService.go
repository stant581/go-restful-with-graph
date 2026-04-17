package services

import (
	"github.com/stant581/go-restful/models"
	"github.com/stant581/go-restful/repositories"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func (service *UserService) GetUserById(id int) (models.User, error) {
	return service.UserRepo.FindById(id)
}

func (service *UserService) CreateUser(user models.User) (models.User, error) {
	return service.UserRepo.Save(user)
}
