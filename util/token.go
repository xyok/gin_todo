package util

import (
	"fmt"
	"gin_todo/setting"
	"github.com/dgrijalva/jwt-go"
)

func ParseToken(tokenString string) (jwt.MapClaims) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(setting.AppSetting.JwtSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	} else {
		return nil
	}

}

func GenToken(claims map[string]interface{}) (tokenString string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   claims["id"],
		"name": claims["name"],
		"exp":  claims["exp"],
	})
	secret := setting.AppSetting.JwtSecret
	tokenString, _ = token.SignedString([]byte(secret))
	return tokenString

}
