package models

type Account struct {
	AccountId   int    `json:"account_id"`
	AccountName string `json:"account_name"`
	Status 	    bool   `json:"status"`
}