package diaries

import (
	"time"

	"github.com/henrioseptiano/taptalk-diary/entity"
	"github.com/henrioseptiano/taptalk-diary/models"
)

type DiariesServicesInterface interface {
	Create(int64, *models.ReqCreateDiary) error
	Update(int64, int64, *models.ReqUpdateDiary) error
	Delete(int64, int64) error
	GetAllDiariesPagination(string, []int, int, int) *models.Pagination
	GetDiariesByID(int64, int64) *entity.UserDiary
}

type DiariesRepositoriesInterface interface {
	Create(*entity.UserDiary) error
	GetDiaryByID(int64, int64) *entity.UserDiary
	GetAllDiariesPagination(string, []int, int, int) (*[]entity.UserDiary, int, int)
	Update(int64, int64, *entity.UserDiary) error
	Delete(int64, int64) error
	CheckDiaryByDateUser(time.Time, int64) int
	GetDiaryByDateUser(time.Time, int64) *entity.UserDiary
}
