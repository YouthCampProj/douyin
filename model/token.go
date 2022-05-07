package model

type Token struct {
	Common
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

// GetToken 获取Token
func GetToken(user *User) string {
	return user.Name
	// TODO: 获取Token的具体实现
}
