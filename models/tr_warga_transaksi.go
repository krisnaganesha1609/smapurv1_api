package models

import (
	"time"
)

type WargaTransaksi struct {
	KD_Warga           string         `gorm:"type:VARCHAR(10);column:warga_id;primaryKey" json:"warga_id"`
	Warga              Warga          `gorm:"foreignKey:KD_Warga;references:ID_Warga" json:"warga"`
	KD_Sandi           string         `gorm:"type:VARCHAR(10);column:kd_sandi;primarykey" json:"kd_sandi"`
	Sandi              SandiTransaksi `gorm:"foreignKey:KD_Sandi;references:ID_Sandi" json:"sandi_transaksi"`
	Tgl_Transaksi      time.Time      `gorm:"type:DATETIME;column:tgl_transaksi;primaryKey" json:"tgl_transaksi"`
	Serial_Number      int            `gorm:"type:INT;column:no_urut;primaryKey" json:"no_urut"`
	TransactionNominal float32        `gorm:"type:NUMERIC(22,2);column:nom_transaksi;" json:"nom_transaksi"`
	Status             string         `gorm:"type:CHAR(1);column:status" json:"status"`
	Created_At         time.Time      `gorm:"type:DATETIME;column:tgl_rekam" json:"tgl_rekam"`
	Creator            string         `gorm:"type:VARCHAR(30);column:petugas_rekam" json:"petugas_rekam"`
	Updated_At         time.Time      `gorm:"type:DATETIME;column:tgl_ubah" json:"tgl_ubah"`
	Updater            string         `gorm:"type:VARCHAR(30);column:petugas_ubah" json:"petugas_ubah"`
}

type CreateWargaTransaksiRequest struct {
	KD_Warga           string    `json:"warga_id" binding:"required"`
	KD_Sandi           string    `json:"kd_sandi" binding:"required"`
	Tgl_Transaksi      time.Time `json:"tgl_transaksi" binding:"required"`
	TransactionNominal float32   `json:"nom_transaksi" binding:"required"`
	Serial_Number      int       `json:"no_urut,omitempty"`
	Status             string    `json:"status,omitempty"`
}

type UpdateWargaTransaksiRequest struct {
	KD_Warga           string    `json:"warga_id,omitempty"`
	KD_Sandi           string    `json:"kd_sandi,omitempty"`
	Tgl_Transaksi      time.Time `json:"tgl_transaksi,omitempty"`
	TransactionNominal float32   `json:"nom_transaksi,omitempty"`
	Serial_Number      int       `json:"no_urut,omitempty"`
	Status             string    `json:"status,omitempty"`
}

func (wt *WargaTransaksi) TableName() string {
	return "tr_warga_transaksi"
}
