package services

import (
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/henrioseptiano/taptalk-diary/models"
)

type UserServices struct {
}

func New() *UserServices {
	return &UserServices{}
}

func (us *UserServices) LoginUser(users *models.ReqUserLogin) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	mySecretKey := []byte(secretKey)
	claims := models.UserClaims{
		35,
		users.Username,
		jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
