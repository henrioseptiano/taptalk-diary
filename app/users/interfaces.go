package users

import (
	"github.com/henrioseptiano/taptalk-diary/entity"
	"github.com/henrioseptiano/taptalk-diary/models"
	"gorm.io/gorm"
)

type UserServicesInterfaces interface {
	LoginUser(*models.ReqUserLogin) (string, error)
	RegisterUser(*models.ReqUserRegister) error
}

type UserRepositoryInterfaces interface {
	InsertUser(*entity.MasterUser) (*gorm.DB, error)
	CreateUserAuth(*gorm.DB, *entity.UserAuth) error
}
