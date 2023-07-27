package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	KD_User       uuid.UUID `gorm:"type:VARCHAR(100);column:kd_user;primaryKey"`
	User          Users     `gorm:"foreignKey:KD_User;references:ID_User"`
	KD_Role       string    `gorm:"type:VARCHAR(10);column:kd_role;primaryKey"`
	Role          Role      `gorm:"foreignKey:KD_Role;references:ID_Role"`
	Default_Role  string    `gorm:"type:VARCHAR(100);column:default_role"`
	Info          string    `gorm:"type:VARCHAR(300);column:keterangan"`
	Status        string    `gorm:"type:CHAR(1);column:status"`
	Serial_Number int       `gorm:"type:INT;column:no_urut"`
	Created_At    time.Time `gorm:"type:DATETIME;column:tgl_rekam"`
	Creator       string    `gorm:"type:VARCHAR(30);column:petugas_rekam"`
	Updated_At    time.Time `gorm:"type:DATETIME;column:tgl_ubah"`
	Updater       string    `gorm:"type:VARCHAR(30);column:petugas_ubah"`
}

type CreateUserWithRoleRequest struct {
	KD_User       uuid.UUID `json:"kd_user" binding:"required"`
	KD_Role       string    `json:"kd_role" binding:"required"`
	Default_Role  string    `json:"default_role" binding:"required"`
	Info          string    `json:"keterangan,omitempty"`
	Status        string    `json:"status,omitempty"`
	Serial_Number int       `json:"no_urut,omitempty"`
}

type UpdateUserWithRoleRequest struct {
	KD_User       uuid.UUID `json:"kd_user,omitempty"`
	KD_Role       string    `json:"kd_role,omitempty"`
	Default_Role  string    `json:"default_role,omitempty"`
	Info          string    `json:"keterangan,omitempty"`
	Status        string    `json:"status,omitempty"`
	Serial_Number int       `json:"no_urut,omitempty"`
}

type UserWithRoleResponse struct {
	KD_User       uuid.UUID `json:"kd_user"`
	User          Users     `json:"user"`
	KD_Role       string    `json:"kd_role"`
	Role          Role      `json:"role"`
	Default_Role  string    `json:"default_role"`
	Info          string    `json:"keterangan,omitempty"`
	Status        string    `json:"status"`
	Serial_Number int       `json:"no_urut"`
	Created_At    time.Time `json:"tgl_rekam,omitempty"`
	Updated_At    time.Time `json:"tgl_ubah,omitempty"`
}

func (ur *UserRole) TableName() string {
	return "ad_user_role"
}
