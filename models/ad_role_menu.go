package models

import "time"

type RoleMenu struct {
	KD_Role        string    `gorm:"type:VARCHAR(10);column:kd_role;primaryKey"`
	Roles          Role      `gorm:"foreignKey:KD_Role;references:ID_Role"`
	KD_Menu        string    `gorm:"type:VARCHAR(10);column:kd_menu;primaryKey"`
	Menus          Menu      `gorm:"foreignKey:KD_Menu;references:ID_Menu"`
	Allow_Add      uint      `gorm:"type:INT(1);column:allow_add"`
	Allow_Edit     uint      `gorm:"type:INT(1);column:allow_edit"`
	Allow_Del      uint      `gorm:"type:INT(1);column:allow_del"`
	Allow_View     uint      `gorm:"type:INT(1);column:allow_view"`
	Allow_Approve  uint      `gorm:"type:INT(1);column:allow_approve"`
	Allow_Download uint      `gorm:"type:INT(1);column:allow_download"`
	Allow_Upload   uint      `gorm:"type:INT(1);column:allow_upload"`
	Visible        uint      `gorm:"type:INT(1);column:st_visible"`
	Group_Role     string    `gorm:"type:VARCHAR(10);column:grup_role"`
	Info           string    `gorm:"type:VARCHAR(300);column:keterangan"`
	Status         string    `gorm:"type:CHAR(1);column:status"`
	Serial_Number  int       `gorm:"type:INT;column:no_urut"`
	Created_At     time.Time `gorm:"type:DATETIME;column:tgl_rekam"`
	Creator        string    `gorm:"type:VARCHAR(30);column:petugas_rekam"`
	Updated_At     time.Time `gorm:"type:DATETIME;column:tgl_ubah"`
	Updater        string    `gorm:"type:VARCHAR(30);column:petugas_ubah"`
}

type CreateRoleMenuRequest struct {
	KD_Role        string `json:"kd_role" binding:"required"`
	KD_Menu        string `json:"kd_menu" binding:"required"`
	Allow_Add      uint   `json:"allow_add,omitempty"`
	Allow_Edit     uint   `json:"allow_edit,omitempty"`
	Allow_Del      uint   `json:"allow_del,omitempty"`
	Allow_View     uint   `json:"allow_view,omitempty"`
	Allow_Approve  uint   `json:"allow_approve,omitempty"`
	Allow_Download uint   `json:"allow_download,omitempty"`
	Allow_Upload   uint   `json:"allow_upload,omitempty"`
	Visible        uint   `json:"st_visible,omitempty"`
	Group_Role     string `json:"grup_role,omitempty"`
	Info           string `json:"keterangan,omitempty"`
	Status         string `json:"status,omitempty"`
	Serial_Number  int    `json:"no_urut,omitempty"`
}

type UpdateRoleMenuRequest struct {
	KD_Role        string `json:"kd_role,omitempty"`
	KD_Menu        string `json:"kd_menu,omitempty"`
	Allow_Add      uint   `json:"allow_add,omitempty"`
	Allow_Edit     uint   `json:"allow_edit,omitempty"`
	Allow_Del      uint   `json:"allow_del,omitempty"`
	Allow_View     uint   `json:"allow_view,omitempty"`
	Allow_Approve  uint   `json:"allow_approve,omitempty"`
	Allow_Download uint   `json:"allow_download,omitempty"`
	Allow_Upload   uint   `json:"allow_upload,omitempty"`
	Visible        uint   `json:"st_visible,omitempty"`
	Group_Role     string `json:"grup_role,omitempty"`
	Info           string `json:"keterangan,omitempty"`
	Status         string `json:"status,omitempty"`
	Serial_Number  int    `json:"no_urut,omitempty"`
}

type RoleMenuResponse struct {
	KD_Role        string    `json:"kd_role"`
	Roles          Role      `json:"role"`
	KD_Menu        string    `json:"kd_menu"`
	Menus          Menu      `json:"menu"`
	Allow_Add      uint      `json:"allow_add"`
	Allow_Edit     uint      `json:"allow_edit"`
	Allow_Del      uint      `json:"allow_del"`
	Allow_View     uint      `json:"allow_view"`
	Allow_Approve  uint      `json:"allow_approve"`
	Allow_Download uint      `json:"allow_download"`
	Allow_Upload   uint      `json:"allow_upload"`
	Visible        uint      `json:"st_visible"`
	Group_Role     string    `json:"grup_role"`
	Info           string    `json:"keterangan"`
	Status         string    `json:"status"`
	Serial_Number  int       `json:"no_urut"`
	Created_At     time.Time `json:"tgl_rekam"`
	Creator        string    `json:"petugas_rekam"`
	Updated_At     time.Time `json:"tgl_ubah"`
	Updater        string    `json:"petugas_ubah"`
}

func (rm *RoleMenu) TableName() string {
	return "ad_role_menu"
}
