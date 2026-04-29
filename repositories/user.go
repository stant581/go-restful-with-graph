package repositories

import (
	"errors"

	"github.com/stant581/go-restful/models"
)

type UserRepository interface {
	FindById(id int) (models.User, error)
	Save(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
}

type inMemoryUserRepository struct {
	users  map[int]models.User
	nextID int
}

func NewUserRepository() UserRepository {
	return &inMemoryUserRepository{
		users:  make(map[int]models.User),
		nextID: 1,
	}
}

func (repo *inMemoryUserRepository) FindById(id int) (models.User, error) {
	user, ok := repo.users[id]
	if !ok {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (repo *inMemoryUserRepository) Save(user models.User) (models.User, error) {
	if user.UserId == 0 {
		user.UserId = repo.nextID
		repo.nextID++
	}

	repo.users[user.UserId] = user
	return user, nil
}

func (repo *inMemoryUserRepository) GetAllUsers() ([]models.User, error) {
	users := []models.User{}
	for _, user := range repo.users {
		users = append(users, user)
	}
	return users, nil
}
