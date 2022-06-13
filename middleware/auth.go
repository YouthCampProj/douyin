package middleware

import (
	"github.com/YouthCampProj/douyin/pkg/auth"
	"github.com/YouthCampProj/douyin/utils"
	"github.com/gin-gonic/gin"
)

const (
	AuthTokenSuccess      = 0
	AuthTokenNotMatchUser = iota + 100
	AuthTokenExpired
	AuthTokenInvalid
	AuthNoTokenProvided
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if res, ok := CheckTokenValidated(token); !ok {
			SendAuthResponse(res, c)
			c.Abort()
			return
		}
		userID := c.Query("user_id")
		if userID != "" && IsTokenMatchUser(token, userID) {
			c.Next()
			return
		}
	}
}

// SendAuthResponse is a helper function to send auth response
func SendAuthResponse(code int, c *gin.Context) {
	switch code {
	case AuthTokenNotMatchUser:
		SendCommonResponse(AuthTokenNotMatchUser, "Token not match user", c)
	case AuthTokenExpired:
		SendCommonResponse(AuthTokenExpired, "Token expired", c)
	case AuthTokenInvalid:
		SendCommonResponse(AuthTokenInvalid, "Token invalid", c)
	case AuthNoTokenProvided:
		SendCommonResponse(AuthNoTokenProvided, "No token provided", c)
	}
}

// IsTokenMatchUser is a helper function to check if token match user or not
// Return true for match, false for not match
func IsTokenMatchUser(token string, userID string) bool {
	user, err := auth.ParseToken(token)
	if err != nil {
		return false
	}
	if utils.Str2uint64(userID) != user.ID {
		// token not match user
		return false
	}
	return true
}

// CheckTokenValidated return 0,true for valid token, AuthResponseCode, false for invalid token
func CheckTokenValidated(tokenString string) (int, bool) {
	// 未提供Token
	if tokenString == "" {
		return AuthNoTokenProvided, false
	}
	// TODO: Token与约束条件不符合
	// 没有懂这个约束条件
	// TODO: Token过期
	_, err := auth.ParseToken(tokenString)
	if err != nil {
		return AuthTokenInvalid, false
	}
	return 0, true
}
