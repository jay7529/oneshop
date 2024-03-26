package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("setting.JwtSecret")

type Claims struct {
	Identity string `json:"identity"`
	ID       int    `json:"id"`
	jwt.StandardClaims
}

// 根據Identity和ID產生token
func GenerateToken(identity string, id int) (string, error) {

	//設置token有效時間
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Hour)
	claims := Claims{
		Identity: identity,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			// 過期時間
			ExpiresAt: expireTime.Unix(),
			// 指定token發行人
			Issuer: "",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
