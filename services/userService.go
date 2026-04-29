package services

import (
	"errors"
	"log"

	"context"

	"github.com/stant581/go-restful/models"
	"github.com/stant581/go-restful/repositories"
	"golang.org/x/sync/errgroup"
)

type UserService struct {
	UserRepo    repositories.UserRepository
	AccountRepo repositories.AccountRepository
}

func (service *UserService) GetUserById(id int) (models.User, error) {
	return service.UserRepo.FindById(id)
}

func (service *UserService) CreateUser(user models.User) (models.User, error) {
	return service.UserRepo.Save(user)
}

func (service *UserService) GetAccountById(id int) (models.Account, error) {
	if id != 0 {
		account, err := service.AccountRepo.FindById(id)
		if err != nil {
			return models.Account{}, err
		}
		return account, nil
	}
	return models.Account{}, errors.New("account not found")
}

func (service *UserService) CreateAccount(account models.Account) (models.Account, error) {
	if account.AccountName == "" {
		return models.Account{}, errors.New("account name is required")
	}
	return service.AccountRepo.Save(account)
}

func (service *UserService) GetUserProfileById(id int) (models.UserProfile, error) {
	user, err1 := service.GetUserById(id)
	if err1 != nil {
		return models.UserProfile{}, err1
	}
	account, err2 := service.GetAccountById(user.AccountId)
	if err2 != nil {
		log.Printf("account not found for user %d: %v", id, err2)
	}
	if account.AccountId == 0 {
		return models.UserProfile{
			User: user,
		}, nil
	}
	return models.UserProfile{
		User:    user,
		Account: account,
	}, nil
}

func (service *UserService) CreateUserProfile(userProfile models.UserProfile) (models.UserProfile, error) {
	createdUserProfile := models.UserProfile{}
	if userProfile.Account.AccountId != 0 {
		account := userProfile.Account
		createdAccount, _ := service.CreateAccount(account)
		createdUserProfile.Account = createdAccount
	}
	if userProfile.User.UserId != 0 {
		user := userProfile.User
		createdUser, err := service.CreateUser(user)
		if err != nil {
			return models.UserProfile{}, err
		}
		createdUserProfile.User = createdUser

	}
	return createdUserProfile, nil
}

func (service *UserService) GetAllUsersAndAccounts() ([]models.UserProfile, error) {
	userProfiles := []models.UserProfile{}

	var users []models.User
	var accounts []models.Account

	g, _ := errgroup.WithContext(context.Background())

	if err := g.Wait(); err != nil {
		return nil, err
	}

	accountMap := make(map[int]models.Account)
	for _, account := range accounts {
		accountMap[account.AccountId] = account
	}

	for _, user := range users {
		userProfile := models.UserProfile{
			User: user,
		}
		if account, ok := accountMap[user.AccountId]; ok {
			userProfile.Account = account
		}
		userProfiles = append(userProfiles, userProfile)
	}
	return userProfiles, nil
}
