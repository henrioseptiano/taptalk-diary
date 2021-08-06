package services

import (
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/henrioseptiano/taptalk-diary/app/users/repository"
	"github.com/henrioseptiano/taptalk-diary/entity"
	"github.com/henrioseptiano/taptalk-diary/models"
	"github.com/henrioseptiano/taptalk-diary/utils"
)

type UserServices struct {
	UserRepository repository.UserRepository
}

func New(userRepository repository.UserRepository) UserServices {
	return UserServices{UserRepository: userRepository}
}

func (us UserServices) LoginUser(users *models.ReqUserLogin) (string, error) {
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

func (us UserServices) RegisterUser(users *models.ReqUserRegister) error {
	masterUsers := &entity.MasterUser{
		Username: users.Username,
		Email:    users.Email,
		Birthday: users.Birthday,
		FullName: users.Fullname,
	}
	trx, err := us.UserRepository.InsertUser(masterUsers)
	if err != nil {
		return err
	}
	userAuth := &entity.UserAuth{
		UserID:   masterUsers.ID,
		Password: utils.HashedPassword(users.Password),
		DeviceID: "",
	}
	err = us.UserRepository.CreateUserAuth(trx, userAuth)
	if err != nil {
		return err
	}
	trx.Commit()
	return nil
}
