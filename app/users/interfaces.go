package users

import "github.com/henrioseptiano/taptalk-diary/models"

type UserServicesInterfaces interface {
	LoginUser(users *models.ReqUserLogin) (string, error)
}
