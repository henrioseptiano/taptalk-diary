package entity

import "time"

type UserDiary struct {
	ID        string     `json:"id" gorm:"primary_key"`
	UserID    int64      `json:"userId" gorm:"not null"`
	Title     string     `json:"title" gorm:"not null"`
	BodyText  string     `json:"bodyText" gorm:"not null"`
	CreatedAt time.Time  `json:"createdAt" gorm:"not null"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
