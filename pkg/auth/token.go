package auth

import (
	"errors"
	"github.com/YouthCampProj/douyin/model"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	TokenExpireDuration = time.Hour * 24 * 365 // jwt过期时间1年
	issuer              = "douyin"             // 签发人
)

var TokenSecret = []byte("douyin") // token盐值

// MyClaims jwt保存信息
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// NewToken 生成新的token
func NewToken(user *model.User) string {
	mc := MyClaims{
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 注册过期时间
			Issuer:    issuer,                                     // 签发人
		},
	}
	// 生成token并返回其字符串
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	tokenString, _ := token.SignedString(TokenSecret)
	// 缺少错误处理
	return tokenString
}

// ParseToken 解析token 获取用户
func ParseToken(tokenString string) (*model.User, error) {
	// 解析token到mc
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return TokenSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 校验token
	if token.Valid {
		return model.GetUserByUsername(mc.Username)
	}
	return nil, errors.New("无效的token")
}

// CheckToken 检查token是否有效
func CheckToken(tokenString string) bool {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return TokenSecret, nil
	})
	return err == nil && token.Valid
}

func CheckTokenWithUserID(token string, userID uint64) bool {
	user, _ := ParseToken(token)
	return user.ID == userID
}
