package entity

import "time"

// MasterUser ...
type MasterUser struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	Username  string     `json:"username" gorm:"unique; not null"`
	Email     string     `json:"email" gorm:"unique; not null"`
	Birthday  string     `json:"birthday" gorm:"not null"`
	FullName  string     `json:"fullName" gorm:"not null"`
	CreatedAt time.Time  `json:"createdAt" gorm:"not null"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
}
