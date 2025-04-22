package jwts

import (
	"fmt"
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
		"token": str,
		"exp":   time.Now().Add(t).Unix(),
	})
	token, err := claims.SignedString([]byte(accessSecret))
	if err != nil {
		zap.L().Error("token signed fail", zap.Error(err))
	}
	refreshClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": str,
		"exp":   time.Now().Add(rf).Unix(),
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

func ParseToken(tokenString string, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method:%v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		zap.L().Error("parseToken errr: ", zap.Error(err))
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		val := claims["token"].(string)
		exp := int64(claims["exp"].(float64))
		if exp <= time.Now().Unix() {
			return "", fmt.Errorf("token过期了")
		}
		return val, nil
	} else {
		return "", err
	}
}
