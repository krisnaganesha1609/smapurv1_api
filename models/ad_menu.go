package models

import "time"

type Menu struct {
	ID_Menu        string    `gorm:"type:VARCHAR(10);column:kd_menu;primaryKey" json:"kd_menu"`
	Menu_Name      string    `gorm:"type:VARCHAR(100);column:nama_menu" json:"nama_menu"`
	KD_Parent_Menu string    `gorm:"type:VARCHAR(10);column:kd_menu_induk" json:"kd_menu_induk"`
	Level_Menu     string    `gorm:"type:VARCHAR(10);column:level_menu" json:"level_menu"`
	Group_Menu     string    `gorm:"type:VARCHAR(10);column:grup_menu" json:"grup_menu"`
	Link           string    `gorm:"type:VARCHAR(300);column:link" json:"link"`
	Icon_Menu      []byte    `gorm:"type:BLOB;column:icon_menu" json:"icon_menu"`
	Info           string    `gorm:"type:VARCHAR(300);column:keterangan" json:"keterangan"`
	Status         string    `gorm:"CHAR(1);column:status" json:"status"`
	Serial_Number  int       `gorm:"type:INT;column:no_urut" json:"no_urut"`
	Created_At     time.Time `gorm:"type:DATETIME;column:tgl_rekam" json:"tgl_rekam"`
	Creator        string    `gorm:"type:VARCHAR(30);column:petugas_rekam" json:"petugas_rekam"`
	Updated_At     time.Time `gorm:"type:DATETIME;column:tgl_ubah" json:"tgl_ubah"`
	Updater        string    `gorm:"type:VARCHAR(30);column:petugas_ubah" json:"petugas_ubah"`
}

type CreateMenuRequest struct {
	ID_Menu        string `json:"kd_menu" binding:"required"`
	Menu_Name      string `json:"nama_menu" binding:"required"`
	KD_Parent_Menu string `json:"kd_menu_induk" binding:"required"`
	Level_Menu     string `json:"level_menu" binding:"required"`
	Group_Menu     string `json:"grup_menu" binding:"required"`
	Link           string `json:"link" binding:"required"`
	Icon_Menu      []byte `json:"icon_menu" binding:"required"`
	Info           string `json:"keterangan,omitempty"`
	Status         string `json:"status,omitempty"`
	Serial_Number  int    `json:"no_urut,omitempty"`
}

type UpdateMenuRequest struct {
	Menu_Name      string `json:"nama_menu,omitempty"`
	KD_Parent_Menu string `json:"kd_menu_induk,omitempty"`
	Level_Menu     string `json:"level_menu,omitempty"`
	Group_Menu     string `json:"grup_menu,omitempty"`
	Link           string `json:"link,omitempty"`
	Icon_Menu      []byte `json:"icon_menu,omitempty"`
	Info           string `json:"keterangan,omitempty"`
	Status         string `json:"status,omitempty"`
	Serial_Number  int    `json:"no_urut,omitempty"`
}

type MenuResponse struct {
	ID_Menu        string    `json:"kd_menu"`
	Menu_Name      string    `json:"nama_menu"`
	KD_Parent_Menu string    `json:"kd_menu_induk,omitempty"`
	Level_Menu     string    `json:"level_menu,omitempty"`
	Group_Menu     string    `json:"grup_menu,omitempty"`
	Link           string    `json:"link"`
	Icon_Menu      []byte    `json:"icon_menu"`
	Info           string    `json:"keterangan,omitempty"`
	Status         string    `json:"status,omitempty"`
	Serial_Number  int       `json:"no_urut,omitempty"`
	Created_At     time.Time `json:"tgl_rekam,omitempty"`
	Creator        string    `json:"petugas_rekam,omitempty"`
	Updated_At     time.Time `json:"tgl_ubah,omitempty"`
	Updater        string    `json:"petugas_ubah,omitempty"`
}

func (m *Menu) TableName() string {
	return "ad_menu"
}
