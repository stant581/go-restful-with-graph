package repositories

import (
	"errors"

	"github.com/stant581/go-restful/models"
)

type AccountRepository interface {
	FindById(id int) (models.Account, error)
	Save(account models.Account) (models.Account, error)
	GetAllAccounts() ([]models.Account, error)
}

type inMemoryAccountRepository struct {
	accounts map[int]models.Account
	nextID   int
}

func NewAccountRepository() AccountRepository {
	return &inMemoryAccountRepository{
		accounts: make(map[int]models.Account),
		nextID:   1,
	}
}

func (repo *inMemoryAccountRepository) FindById(id int) (models.Account, error) {
	account, ok := repo.accounts[id]
	if !ok {
		return models.Account{}, errors.New("account not found")
	}
	return account, nil
}

func (repo *inMemoryAccountRepository) Save(account models.Account) (models.Account, error) {
	if account.AccountId == 0 {
		account.AccountId = repo.nextID
		repo.nextID++
	}
	repo.accounts[account.AccountId] = account
	return account, nil
}

func (repo *inMemoryAccountRepository) GetAllAccounts() ([]models.Account, error) {
	accounts := []models.Account{}
	for _, account := range repo.accounts {
		accounts = append(accounts, account)
	}
	return accounts, nil
}
