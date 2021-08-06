package repository

import (
	"log"
	"time"

	"github.com/henrioseptiano/taptalk-diary/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (ur UserRepository) InsertUser(users *entity.MasterUser) (*gorm.DB, error) {
	trx := ur.DB.Begin()
	if err := trx.Create(users).Error; err != nil {
		trx.Rollback()
		log.Printf("ERROR: %s", err.Error())
		return trx, err
	}
	log.Printf("Created user: %+v\n", users)
	return trx, nil
}

func (ur UserRepository) CreateUserAuth(trx *gorm.DB, userAuth *entity.UserAuth) error {
	if err := trx.Create(userAuth).Error; err != nil {
		trx.Rollback()
		log.Printf("ERROR: %s", err.Error())
		return err
	}
	log.Printf("Created user: %+v\n", userAuth)
	return nil
}

func (ur UserRepository) GetCurrentDeviceID(userID int64) string {
	var foundedDeviceId string
	ur.DB.Where("user_id = ?", userID).Select("device_id").Table("user_auths").Scan(&foundedDeviceId)
	return foundedDeviceId
}

func (ur UserRepository) GetUserAuth(userID int64) *entity.UserAuth {
	userAuth := new(entity.UserAuth)
	ur.DB.Where("user_id = ?", userID).Find(userAuth)
	if userAuth == nil {
		return nil
	}
	return userAuth
}

func (ur UserRepository) GetUserByUsername(username string) *entity.MasterUser {
	masterUser := new(entity.MasterUser)
	ur.DB.Where("username = ? OR email = ?", username, username).Find(masterUser)
	if masterUser == nil {
		return nil
	}
	return masterUser
}

func (ur UserRepository) UpdateDeviceIDLastLogin(deviceID string, userID int64) error {
	updateData := map[string]interface{}{"device_id": deviceID, "last_login": time.Now()}
	if err := ur.DB.Model(&entity.UserAuth{}).Where("user_id = ?", userID).Updates(updateData).Error; err != nil {
		log.Printf("ERROR: %s", err.Error())
		return err
	}
	return nil
}
