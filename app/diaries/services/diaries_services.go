package services

import (
	"errors"
	"time"

	"github.com/henrioseptiano/taptalk-diary/app/diaries/repository"
	"github.com/henrioseptiano/taptalk-diary/entity"
	"github.com/henrioseptiano/taptalk-diary/models"
	"github.com/henrioseptiano/taptalk-diary/utils"
)

type DiariesServices struct {
	DiariesRepositories repository.DiariesRepository
}

func New(diariesRepository repository.DiariesRepository) DiariesServices {
	return DiariesServices{DiariesRepositories: diariesRepository}
}

func (ds DiariesServices) Create(UserID int64, modelCreateDiary *models.ReqCreateDiary) error {
	datePost, err := time.Parse("02-01-2006", modelCreateDiary.DatePost)
	if err != nil {
		return errors.New("Invalid Date Post")
	}

	count := ds.DiariesRepositories.CheckDiaryByDateUser(datePost, UserID)
	if count > 0 {
		entityDiary := ds.DiariesRepositories.GetDiaryByDateUser(datePost, UserID)
		entityDiary.Title = modelCreateDiary.Title
		entityDiary.BodyText = modelCreateDiary.BodyText
		if err := ds.DiariesRepositories.Update(entityDiary.ID, UserID, entityDiary); err != nil {
			return err
		}
		return nil
	}

	entityDiary := entity.UserDiary{
		UserID:   UserID,
		DatePost: datePost,
		Title:    modelCreateDiary.Title,
		BodyText: modelCreateDiary.BodyText,
	}

	if err := ds.DiariesRepositories.Create(&entityDiary); err != nil {
		return err
	}
	return nil
}

func (ds DiariesServices) GetAllDiariesPagination(year string, monthRange []int, page, limit int) *models.Pagination {
	data, count, offset := ds.DiariesRepositories.GetAllDiariesPagination(year, monthRange, page, limit)
	paginatedData := &models.Pagination{
		TotalRecords: count,
		TotalPages:   utils.TotalPages(count, limit),
		Data:         data,
		Offset:       offset,
		Limit:        limit,
		Page:         page,
		PrevPage:     utils.PrevPage(page),
		NextPage:     utils.NextPage(page, utils.TotalPages(count, limit)),
	}
	return paginatedData
}

func (ds DiariesServices) GetDiariesByID(id int64, userID int64) *entity.UserDiary {
	return ds.DiariesRepositories.GetDiaryByID(id, userID)
}

func (ds DiariesServices) Update(UserID int64, ID int64, modelUpdatediary *models.ReqUpdateDiary) error {
	datePost, err := time.Parse("02-01-2006", modelUpdatediary.DatePost)
	if err != nil {
		return errors.New("Invalid Date Post")
	}

	entityDiary := ds.DiariesRepositories.GetDiaryByID(ID, UserID)
	entityDiary.DatePost = datePost
	entityDiary.Title = modelUpdatediary.Title
	entityDiary.BodyText = modelUpdatediary.BodyText
	if err := ds.DiariesRepositories.Update(ID, UserID, entityDiary); err != nil {
		return err
	}
	return nil
}

func (ds DiariesServices) Delete(UserID, ID int64) error {
	return ds.DiariesRepositories.Delete(ID, UserID)
}
