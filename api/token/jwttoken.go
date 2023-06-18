package token

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT(username string) (string, error) {

	var mySigningKey = []byte(jwtKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		//fmt.Errorf("Something went wrong which signing token: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
