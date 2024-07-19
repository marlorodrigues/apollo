package authentication

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func validToken(tokenString string) bool {

	// Verificar se o token está presente
	if tokenString == "" {
		fmt.Println("Token was empty")
		return false
	}

	jwtKey = []byte("db0094286db986ab65a6c6621e9f49a15c41bf3e857a681bfd574a54ed7f4077")

	// Validar o token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Token was not recognized!")
			return nil, fmt.Errorf("token was not recognized")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		fmt.Println("Token was not valid!", err.Error())
		return false
	}

	// Passar para o próximo middleware
	return true
}
