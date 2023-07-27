package models

import "time"

type Agama struct {
	ID_Agama      string    `gorm:"type:VARCHAR(10);column:kd_agama;primaryKey" json:"kd_agama"`
	Nama_Agama    string    `gorm:"type:VARCHAR(100);column:nama_agama" json:"nama_agama"`
	Info          string    `gorm:"type:VARCHAR(300);column:keterangan" json:"keterangan,omitempty"`
	Status        string    `gorm:"type:CHAR(1);column:status" json:"status,omitempty"`
	Serial_Number int       `gorm:"type:INT;column:no_urut" json:"no_urut,omitempty"`
	Created_At    time.Time `gorm:"type:DATETIME;column:tgl_rekam" json:"tgl_rekam,omitempty"`
	Creator       string    `gorm:"type:VARCHAR(30);column:petugas_rekam" json:"petugas_rekam,omitempty"`
	Updated_At    time.Time `gorm:"type:DATETIME;column:tgl_ubah" json:"tgl_ubah,omitempty"`
	Updater       string    `gorm:"type:VARCHAR(30);column:petugas_ubah" json:"petugas_ubah,omitempty"`
}

type CreateAgamaRequest struct {
	ID_Agama      string `json:"kd_agama" binding:"required"`
	Nama_Agama    string `json:"nama_agama" binding:"required"`
	Info          string `json:"keterangan,omitempty"`
	Status        string `json:"status,omitempty"`
	Serial_Number int    `json:"no_urut,omitempty"`
}

type UpdateAgamaRequest struct {
	Nama_Agama    string `json:"nama_agama,omitempty"`
	Info          string `json:"keterangan,omitempty"`
	Status        string `json:"status,omitempty"`
	Serial_Number int    `json:"no_urut,omitempty"`
}

func (a *Agama) TableName() string {
	return "ms_agama"
}
