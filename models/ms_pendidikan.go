package models

import "time"

type Pendidikan struct {
	ID_Pendidikan   string    `gorm:"type:VARCHAR(10);column:kd_pendidikan;primaryKey" json:"kd_pendidikan"`
	Nama_Pendidikan string    `gorm:"type:VARCHAR(100);column:nama_pendidikan" json:"nama_pendidikan"`
	Info            string    `gorm:"type:VARCHAR(300);column:keterangan" json:"keterangan"`
	Status          string    `gorm:"type:CHAR(1);column:status" json:"status"`
	Serial_Number   int       `gorm:"type:INT;column:no_urut" json:"no_urut"`
	Created_At      time.Time `gorm:"type:DATETIME;column:tgl_rekam" json:"tgl_rekam"`
	Creator         string    `gorm:"type:VARCHAR(30);column:petugas_rekam" json:"petugas_rekam"`
	Updated_At      time.Time `gorm:"type:DATETIME;column:tgl_ubah" json:"tgl_ubah"`
	Updater         string    `gorm:"type:VARCHAR(30);column:petugas_ubah" json:"petugas_ubah"`
}

type CreatePendidikanRequest struct {
	ID_Pendidikan   string `json:"kd_pendidikan" binding:"required"`
	Nama_Pendidikan string `json:"nama_pendidikan" binding:"required"`
	Info            string `json:"keterangan,omitempty"`
	Status          string `json:"status,omitempty"`
	Serial_Number   int    `json:"no_urut,omitempty"`
}

type UpdatePendidikanRequest struct {
	Nama_Pendidikan string `json:"nama_pendidikan"`
	Info            string `json:"keterangan,omitempty"`
	Status          string `json:"status,omitempty"`
	Serial_Number   int    `json:"no_urut,omitempty"`
}

func (p *Pendidikan) TableName() string {
	return "ms_pendidikan"
}
