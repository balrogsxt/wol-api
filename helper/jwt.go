package helper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"wolapi/app"
)

func JwtBuild(_map jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, _map)
	jwt, err := token.SignedString([]byte(app.Config.JwtKey))
	return jwt, err
}
func JwtParse(str string) (jwt.MapClaims, error) {
	claim, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.Config.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	r, flag := claim.Claims.(jwt.MapClaims)
	if !flag {
		return nil, errors.New("解析失败")
	}
	return r, nil
}

