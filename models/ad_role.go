package models

import "time"

type Role struct {
	ID_Role       string    `gorm:"type:VARCHAR(10);column:kd_role;primaryKey"`
	Role_Name     string    `gorm:"type:VARCHAR(100);column:nama_role"`
	Info          string    `gorm:"type:VARCHAR(300);column:keterangan"`
	Icon_Role     []byte    `gorm:"type:BLOB;column:icon_role"`
	Group_Role    string    `gorm:"type:VARCHAR(100);column:grup_role"`
	Status        string    `gorm:"type:CHAR(1);column:status;default:1"`
	Serial_Number int       `gorm:"type:INT;column:no_urut"`
	Created_At    time.Time `gorm:"type:DATETIME;column:tgl_rekam"`
	Creator       string    `gorm:"type:VARCHAR(30);column:petugas_rekam"`
	Updated_At    time.Time `gorm:"type:DATETIME;column:tgl_ubah"`
	Updater       string    `gorm:"type:VARCHAR(30);column:petugas_ubah"`
}

type CreateRoleRequest struct {
	ID_Role       string `json:"kd_role" binding:"required"`
	Role_Name     string `json:"nama_role" binding:"required"`
	Info          string `json:"keterangan,omitempty"`
	Icon_Role     []byte `json:"icon_role" binding:"required"`
	Group_Role    string `json:"group_role" binding:"required"`
	Serial_Number int    `json:"no_urut,omitempty"`
}

type UpdateRoleRequest struct {
	Role_Name     string `json:"nama_role,omitempty"`
	Info          string `json:"keterangan,omitempty"`
	Icon_Role     []byte `json:"icon_role,omitempty"`
	Group_Role    string `json:"group_role,omitempty"`
	Serial_Number int    `json:"no_urut,omitempty"`
}

type RoleResponse struct {
	ID_Role       string    `json:"kd_role"`
	Role_Name     string    `json:"nama_role"`
	Info          string    `json:"keterangan,omitempty"`
	Icon_Role     []byte    `json:"icon_role"`
	Group_Role    string    `json:"group_role"`
	Status        string    `json:"status,omitempty"`
	Serial_Number int       `json:"no_urut,omitempty"`
	Created_At    time.Time `json:"tgl_rekam,omitempty"`
	Creator       string    `json:"petugas_rekam,omitempty"`
	Updated_At    time.Time `json:"tgl_ubah,omitempty"`
	Updater       string    `json:"petugas_ubah,omitempty"`
}

func (r *Role) TableName() string {
	return "ad_role"
}
