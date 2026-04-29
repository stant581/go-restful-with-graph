package models

type UserProfile struct {
	User    `json:"user"`
	Account `json:"account"`
}