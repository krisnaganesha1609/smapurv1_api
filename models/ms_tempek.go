package models

import "time"

type Tempek struct {
	ID_Tempek     string    `gorm:"type:VARCHAR(10);column:kd_tempek;primaryKey" json:"kd_tempek"`
	Nama_Tempek   string    `gorm:"type:VARCHAR(100);column:nama_tempek" json:"nama_tempek"`
	Info          string    `gorm:"type:VARCHAR(300);column:keterangan" json:"keterangan"`
	Status        string    `gorm:"type:CHAR(1);column:status" json:"status"`
	Serial_Number int       `gorm:"type:INT;column:no_urut" json:"no_urut"`
	Created_At    time.Time `gorm:"type:DATETIME;column:tgl_rekam" json:"tgl_rekam"`
	Creator       string    `gorm:"type:VARCHAR(30);column:petugas_rekam" json:"petugas_rekam"`
	Updated_At    time.Time `gorm:"type:DATETIME;column:tgl_ubah" json:"tgl_ubah"`
	Updater       string    `gorm:"type:VARCHAR(30);column:petugas_ubah" json:"petugas_ubah"`
}

type CreateTempekRequest struct {
	ID_Tempek     string `json:"kd_tempek" binding:"required"`
	Nama_Tempek   string `json:"nama_tempek" binding:"required"`
	Info          string `json:"keterangan,omitempty"`
	Status        string `json:"status,omitempty"`
	Serial_Number int    `json:"no_urut,omitempty"`
}

type UpdateTempekRequest struct {
	Nama_Tempek   string `json:"nama_tempek"`
	Info          string `json:"keterangan,omitempty"`
	Status        string `json:"status,omitempty"`
	Serial_Number int    `json:"no_urut,omitempty"`
}

func (t *Tempek) TableName() string {
	return "ms_tempek"
}
