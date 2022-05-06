package model

type Token struct {
	Common
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}
