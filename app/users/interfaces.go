package users

import (
	"github.com/henrioseptiano/taptalk-diary/entity"
	"github.com/henrioseptiano/taptalk-diary/models"
	"gorm.io/gorm"
)

type UserServicesInterfaces interface {
	LoginUser(*models.ReqUserLogin) (string, error)
	RegisterUser(*models.ReqUserRegister) error
	GetCurrentDeviceID(int64) string
}

type UserRepositoryInterfaces interface {
	InsertUser(*entity.MasterUser) (*gorm.DB, error)
	CreateUserAuth(*gorm.DB, *entity.UserAuth) error
	GetCurrentDeviceID(int64) string
	GetUserAuth(int64) *entity.UserAuth
	GetUserByUsername(string) *entity.MasterUser
	UpdateDeviceIDLastLogin(string, int64) error
}
