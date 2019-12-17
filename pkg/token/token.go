package token

import (
	"github.com/dgrijalva/jwt-go"
	"gomall/pkg/settings"
	"gomall/pkg/util"
	"time"
)

type Claims struct {
	Domain   string `json:"domain"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var secret []byte

func Setup() {
	secret = []byte(settings.AppConfig.Auth.JwtSecret)
}

func Generate(domain, username, password string) (string, error) {
	now := time.Now()
	expireTime := now.Add(settings.AppConfig.Auth.ExpireTime)
	claims := Claims{
		util.EncodeMD5(domain),
		util.EncodeMD5(username),
		util.EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    domain,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(secret)
	return token, err
}

func Parse(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
