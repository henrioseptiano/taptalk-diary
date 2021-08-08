package repository

import (
	"log"
	"strconv"
	"time"

	"github.com/henrioseptiano/taptalk-diary/entity"
	"github.com/henrioseptiano/taptalk-diary/utils"
	"gorm.io/gorm"
)

type DiariesRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) DiariesRepository {
	return DiariesRepository{db}
}

func (ds DiariesRepository) Create(entityCreateDiary *entity.UserDiary) error {
	if err := ds.DB.Create(&entityCreateDiary).Error; err != nil {
		log.Printf("ERROR: %s", err.Error())
		return err
	}
	return nil
}

func (ds DiariesRepository) Update(ID int64, UserID int64, entityUpdateDiary *entity.UserDiary) error {
	if err := ds.DB.Model(&entity.UserDiary{}).Where("id = ? AND user_id = ?", ID, UserID).Updates(&entityUpdateDiary).Error; err != nil {
		log.Printf("ERROR: %s", err.Error())
		return err
	}
	return nil
}

func (ds DiariesRepository) GetDiaryByID(ID int64, UserID int64) *entity.UserDiary {
	userDiary := new(entity.UserDiary)
	ds.DB.Where("id = ? AND user_id = ?", ID, UserID).Find(userDiary)
	if userDiary == nil {
		return nil
	}
	return userDiary
}

func (ds DiariesRepository) CheckDiaryByDateUser(date time.Time, userID int64) int {
	var count int
	dateConvert := date.Format("2006-01-02")
	ds.DB.Where("date_post = ? AND user_id = ?", dateConvert, userID).Table("user_diaries").Select("COUNT(*) as count").Scan(&count)
	return count
}

func (ds DiariesRepository) GetDiaryByDateUser(date time.Time, userID int64) *entity.UserDiary {
	userDiary := new(entity.UserDiary)
	dateConvert := date.Format("2006-01-02")
	ds.DB.Model(&entity.UserDiary{}).Where("date_post = ? AND user_id = ?", dateConvert, userID).Find(&userDiary)
	return userDiary
}

func (ds DiariesRepository) Delete(ID int64, UserID int64) error {
	if err := ds.DB.Where("id = ? AND User_id = ?", ID, UserID).Delete(&entity.UserDiary{}).Error; err != nil {
		log.Printf("ERROR: %s", err.Error())
		return err
	}
	return nil
}

func (ds DiariesRepository) GetAllDiariesPagination(year string, monthRange []int, page, limit int) (*[]entity.UserDiary, int, int) {
	userDiaries := new([]entity.UserDiary)
	var count int64 = 0
	yearConv, _ := strconv.Atoi(year)
	offset := utils.Offset(page, limit)
	qry := ds.DB.Model(&entity.UserDiary{}).Where("DATE_FORMAT(date_post, '%Y') = ? AND DATE_FORMAT(date_post, '%m') >= ? AND DATE_FORMAT(date_post, '%m') <= ?",
		yearConv, monthRange[0], monthRange[1]).Offset(offset).Limit(limit)
	qry = qry.Order("date_post ASC").Scan(&userDiaries)
	qry = qry.Offset(0).Count(&count)
	return userDiaries, int(count), offset
}
