package repositories

import (
	"errors"

	"github.com/stant581/go-restful/models"
)

type UserRepository struct {
	users  map[int]models.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make(map[int]models.User),
		nextID: 1,
	}
}

func (repo *UserRepository) FindById(id int) (models.User, error) {
	user, ok := repo.users[id]
	if !ok {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (repo *UserRepository) Save(user models.User) (models.User, error) {
	if user.UserId == 0 {
		user.UserId = repo.nextID
		repo.nextID++
	}

	repo.users[user.UserId] = user
	return user, nil
}
