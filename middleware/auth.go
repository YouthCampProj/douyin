package middleware

import (
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
	// TODO: 对于传入了userID的请求 需要校验其是否与Token一致
	return true
}

// CheckTokenValidated return 0,true for valid token, AuthResponseCode, false for invalid token
func CheckTokenValidated(token string) (int, bool) {
	// 未提供Token
	if token == "" {
		return AuthNoTokenProvided, false
	}
	// TODO: Token与约束条件不符合
	// TODO: Token过期

	return 0, true
}
