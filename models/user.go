package models

type User struct {
	UserId    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	AccountId int    `json:"account_id"`
}
