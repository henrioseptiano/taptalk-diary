package models

import "github.com/golang-jwt/jwt"

// ReqUserLogin ...
type ReqUserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	DeviceID string `json:"deviceID" binding:"required"`
}

// ReqUserRegister ...
type ReqUserRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}

// UserClaims ...
type UserClaims struct {
	ID       int64  `json:"userID"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// JwtCustomClaims
type JwtCustomClaims struct {
	ID       int64
	Username string
}
