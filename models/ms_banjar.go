package models

import "time"

type Banjar struct {
	ID_Banjar     string    `gorm:"type:VARCHAR(10);column:kd_banjar;primaryKey" json:"kd_banjar"`
	Nama_Banjar   string    `gorm:"type:VARCHAR(100);column:nama_banjar" json:"nama_banjar"`
	Info          string    `gorm:"type:VARCHAR(300);column:keterangan" json:"keterangan,omitempty"`
	Status        string    `gorm:"type:CHAR(1);column:status" json:"status,omitempty"`
	Serial_Number int       `gorm:"type:INT;column:no_urut" json:"no_urut,omitempty"`
	Created_At    time.Time `gorm:"type:DATETIME;column:tgl_rekam" json:"tgl_rekam,omitempty"`
	Creator       string    `gorm:"type:VARCHAR(30);column:petugas_rekam" json:"petugas_rekam,omitempty"`
	Updated_At    time.Time `gorm:"type:DATETIME;column:tgl_ubah" json:"tgl_ubah,omitempty"`
	Updater       string    `gorm:"type:VARCHAR(30);column:petugas_ubah" json:"petugas_ubah,omitempty"`
}

type CreateBanjarRequest struct {
	ID_Banjar     string `json:"kd_banjar" binding:"required"`
	Nama_Banjar   string `json:"nama_banjar" binding:"required"`
	Info          string `json:"keterangan,omitempty"`
	Status        string `json:"status,omitempty"`
	Serial_Number int    `json:"no_urut,omitempty"`
}

type UpdateBanjarRequest struct {
	Nama_Banjar   string `json:"nama_banjar,omitempty"`
	Info          string `json:"keterangan,omitempty"`
	Status        string `json:"status,omitempty"`
	Serial_Number int    `json:"no_urut,omitempty"`
}

func (b *Banjar) TableName() string {
	return "ms_banjar"
}
