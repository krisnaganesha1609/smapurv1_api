package models

import "time"

type HubKeluarga struct {
	ID_Hubungan   string    `gorm:"type:VARCHAR(10);column:kd_hub_keluarga;primaryKey" json:"kd_hub_keluarga"`
	Nama_Hubungan string    `gorm:"type:VARCHAR(100);column:nama_hubungan" json:"nama_hubungan"`
	Info          string    `gorm:"type:VARCHAR(300);column:keterangan" json:"keterangan"`
	Status        string    `gorm:"type:CHAR(1);column:status" json:"status"`
	Serial_Number int       `gorm:"type:INT;column:no_urut" json:"no_urut"`
	Created_At    time.Time `gorm:"type:DATETIME;column:tgl_rekam" json:"tgl_rekam"`
	Creator       string    `gorm:"type:VARCHAR(30);column:petugas_rekam" json:"petugas_rekam"`
	Updated_At    time.Time `gorm:"type:DATETIME;column:tgl_ubah" json:"tgl_ubah"`
	Updater       string    `gorm:"type:VARCHAR(30);column:petugas_ubah" json:"petugas_ubah"`
}

type CreateHubKeluargaRequest struct {
	ID_Hubungan   string `json:"kd_hub_keluarga" binding:"required"`
	Nama_Hubungan string `json:"nama_hubungan" binding:"required"`
	Info          string `json:"keterangan,omitempty"`
	Status        string `json:"status,omitempty"`
	Serial_Number int    `json:"no_urut,omitempty"`
}

type UpdateHubKeluargaRequest struct {
	Nama_Hubungan string `json:"nama_hubungan"`
	Info          string `json:"keterangan,omitempty"`
	Status        string `json:"status,omitempty"`
	Serial_Number int    `json:"no_urut,omitempty"`
}

func (hk *HubKeluarga) TableName() string {
	return "ms_hub_keluarga"
}
