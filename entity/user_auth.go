package entity

import "time"

type UserAuth struct {
	UserID    int64      `json:"userId" gorm:"primary_key"`
	Password  string     `json:"password" gorm:"not null"`
	DeviceID  string     `json:"deviceId"`
	LastLogin *time.Time `json:"lastLogin"`
}
