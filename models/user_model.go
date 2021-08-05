package models

// ReqUserLogin ...
type ReqUserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ReqUserRegister ...
type ReqUserRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}
