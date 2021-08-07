package services

import (
	"errors"

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

func (us UserServices) GetCurrentDeviceID(userID int64) string {
	return us.UserRepository.GetCurrentDeviceID(userID)
}

func (us UserServices) LoginUser(users *models.ReqUserLogin) (string, error) {
	getUser := us.UserRepository.GetUserByUsername(users.Username)
	if getUser == nil {
		return "", errors.New("Username or Email is not valid. please register")
	}
	getUserAuth := us.UserRepository.GetUserAuth(getUser.ID)
	if getUserAuth == nil {
		return "", errors.New("Username or Email is not valid. please register")
	}

	if utils.CheckPasswordHash(users.Password, getUserAuth.Password) == false {
		return "", errors.New("Incorrect Username or Password")
	}

	us.UserRepository.UpdateDeviceIDLastLogin(users.DeviceID, getUser.ID)

	tokenString, err := utils.CreateToken(getUser)
	return tokenString, err
}

func (us UserServices) RegisterUser(users *models.ReqUserRegister) error {
	if utils.IsEmailValid(users.Email) == false {
		return errors.New("Email is not valid")
	}
	if utils.IsBirthdayValid(users.Birthday) == false {
		return errors.New("Birthday format is not valid")
	}
	if utils.IsPasswordValid(users.Password) == false {
		return errors.New("Password format must have at least one uppercase, one lowercase, one number, one special character")
	}
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
