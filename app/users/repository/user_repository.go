package repository

import (
	"log"

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
