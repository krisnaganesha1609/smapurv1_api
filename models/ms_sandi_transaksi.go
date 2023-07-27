package models

import "time"

type SandiTransaksi struct {
	ID_Sandi      string    `gorm:"type:VARCHAR(10);column:kd_sandi;primaryKey" json:"kd_sandi"`
	Nama_Sandi    string    `gorm:"type:VARCHAR(100);column:nama_sandi" json:"nama_sandi"`
	Info          string    `gorm:"type:VARCHAR(300);column:keterangan" json:"keterangan"`
	Status        string    `gorm:"type:CHAR(1);column:status" json:"status"`
	Serial_Number int       `gorm:"type:INT;column:no_urut" json:"no_urut"`
	Created_At    time.Time `gorm:"type:DATETIME;column:tgl_rekam" json:"tgl_rekam"`
	Creator       string    `gorm:"type:VARCHAR(30);column:petugas_rekam" json:"petugas_rekam"`
	Updated_At    time.Time `gorm:"type:DATETIME;column:tgl_ubah" json:"tgl_ubah"`
	Updater       string    `gorm:"type:VARCHAR(30);column:petugas_ubah" json:"petugas_ubah"`
}

type CreateSandiTransaksiRequest struct {
	ID_Sandi      string `json:"kd_sandi" binding:"required"`
	Nama_Sandi    string `json:"nama_sandi" binding:"required"`
	Info          string `json:"keterangan,omitempty"`
	Status        string `json:"status,omitempty"`
	Serial_Number int    `json:"no_urut,omitempty"`
}

type UpdateSandiTransaksiRequest struct {
	Nama_Sandi    string `json:"nama_sandi"`
	Info          string `json:"keterangan,omitempty"`
	Status        string `json:"status,omitempty"`
	Serial_Number int    `json:"no_urut,omitempty"`
}

func (st *SandiTransaksi) TableName() string {
	return "ms_sandi_transaksi"
}
