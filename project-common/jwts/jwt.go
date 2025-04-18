package jwts

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type JwtToken struct {
	AccessToken  string
	RefreshToken string
	AccessExp    int64
	RefreshExp   int64
}

// CreateToken 通过 jwt.MapClaims创建token,refreshToken
func CreateToken(str string, accessSecret, refreshSecret string, t time.Duration, rf time.Duration) *JwtToken {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"tokenKey": str,
		"exp":      time.Now().Add(t).Unix(),
	})
	token, err := claims.SignedString([]byte(accessSecret))
	if err != nil {
		zap.L().Error("token signed fail", zap.Error(err))
	}
	refreshClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"key": str,
		"exp": time.Now().Add(rf).Unix(),
	})

	refreshToken, err := refreshClaims.SignedString([]byte(refreshSecret))
	if err != nil {
		zap.L().Error("refresh token signed fail", zap.Error(err))
	}
	return &JwtToken{
		AccessToken:  token,
		RefreshToken: refreshToken,
		AccessExp:    time.Now().Add(t).Unix(),
		RefreshExp:   time.Now().Add(rf).Unix(),
	}

}
