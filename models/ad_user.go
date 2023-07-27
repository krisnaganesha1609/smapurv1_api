package models

import (
	"time"

	"github.com/google/uuid"
)

// Create ad_user table in database
type Users struct {
	ID_User       uuid.UUID `gorm:"type:VARCHAR(100);column:kd_user;primaryKey"`
	Fullname      string    `gorm:"type:VARCHAR(100);column:nama_lengkap"`
	Username      string    `gorm:"type:VARCHAR(100);column:nama_user"`
	Last_Login    time.Time `gorm:"type:DATETIME;column:last_login"`
	Last_Logout   time.Time `gorm:"type:DATETIME;column:last_logout"`
	Photo         []byte    `gorm:"type:BLOB;column:photo"`
	Email         string    `gorm:"type:VARCHAR(100);column:email"`
	NIK           string    `gorm:"type:VARCHAR(100);column:nik"`
	Password      string    `gorm:"type:VARCHAR(100);column:password"`
	Info          string    `gorm:"type:VARCHAR(300);column:keterangan"`
	Status        string    `gorm:"type:CHAR(1);column:status"`
	Serial_Number int       `gorm:"type:INT;column:no_urut"`
	Created_At    time.Time `gorm:"type:DATETIME;column:tgl_rekam"`
	Creator       string    `gorm:"type:VARCHAR(30);column:petugas_rekam"`
	Updated_At    time.Time `gorm:"type:DATETIME;column:tgl_ubah"`
	Updater       string    `gorm:"type:VARCHAR(30);column:petugas_ubah"`
}

func (u *Users) TableName() string {
	return "ad_user"
}

//Create username login request

type UsernameLoginRequest struct {
	Username string `json:"nama_user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//Create NIK login request

type NIKLoginRequest struct {
	NIK      string `json:"nik" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//Create Password verification request

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}

type UpdatePasswordRequest struct {
	Password         string `json:"password" binding:"required"`
	Password_Confirm string `json:"password_confirm" binding:"required"`
}

type CreateUserRequest struct {
	ID_User          uuid.UUID `json:"kd_user" binding:"required"`
	Fullname         string    `json:"nama_lengkap" binding:"required"`
	Username         string    `json:"nama_user" binding:"required"`
	Photo            []byte    `json:"photo" binding:"required"`
	Email            string    `json:"email" binding:"required"`
	NIK              string    `json:"nik" binding:"required"`
	Password         string    `json:"password" binding:"required"`
	Password_Confirm string    `json:"password_confirm" binding:"required"`
	Info             string    `json:"keterangan,omitempty"`
	Status           string    `json:"status,omitempty"`
	Serial_Number    int       `json:"no_urut,omitempty"`
}

type UpdateUserRequest struct {
	Fullname      string `json:"nama_lengkap,omitempty"`
	Username      string `json:"nama_user,omitempty"`
	Photo         []byte `json:"photo,omitempty"`
	Email         string `json:"email,omitempty"`
	NIK           string `json:"nik,omitempty"`
	Info          string `json:"keterangan,omitempty"`
	Status        string `json:"status,omitempty"`
	Serial_Number int    `json:"no_urut,omitempty"`
}

//Create the response struct after sending a request

type UserResponse struct {
	ID_User       uuid.UUID `json:"kd_user,omitempty"`
	Fullname      string    `json:"nama_lengkap,omitempty"`
	Username      string    `json:"nama_user,omitempty"`
	Last_Login    time.Time `json:"last_login,omitempty"`
	Last_Logout   time.Time `json:"last_logout,omitempty"`
	Photo         []byte    `json:"photo,omitempty"`
	Email         string    `json:"email,omitempty"`
	NIK           string    `json:"nik,omitempty"`
	Info          string    `json:"keterangan,omitempty"`
	Status        string    `json:"status,omitempty"`
	Serial_Number int       `json:"no_urut,omitempty"`
	Creator       string    `json:"petugas_rekam,omitempty"`
	Updater       string    `json:"petugas_ubah,omitempty"`
}
