package auth

import "github.com/YouthCampProj/douyin/model"

// NewToken 生成新的token
func NewToken(user *model.User) string {
	return user.Name
	// TODO: 需要生成Token的具体实现
}

// ParseToken 解析token 获取用户
func ParseToken(token string) (*model.User, error) {
	return model.GetUserByUsername(token)
	// TODO: 需要解析Token的具体实现
}

// CheckToken 检查token是否有效
func CheckToken(token string) bool {
	_, err := model.GetUserByUsername(token)
	return err == nil

	// TODO: 需要验证Token的具体实现
}

func CheckTokenWithUserID(token string, userID uint64) bool {
	user, _ := ParseToken(token)
	return user.ID == userID
}
