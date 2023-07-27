package models

import "time"

type WargaKK struct {
	KD_Warga      string    `gorm:"type:VARCHAR(10);column:warga_id;primaryKey" json:"warga_id"`
	Warga         Warga     `gorm:"foreignKey:KD_Warga;references:ID_Warga" json:"warga"`
	NIK           string    `gorm:"type:VARCHAR(100);column:nik;primaryKey" json:"nik"`
	Nama_Warga    string    `gorm:"type:VARCHAR(100);column:nama_warga" json:"nama_warga"`
	Tempat_Lahir  string    `gorm:"type:VARCHAR(100);column:tempat_lahir" json:"tempat_lahir"`
	Tgl_Lahir     time.Time `gorm:"type:DATETIME;column:tgl_lahir" json:"tgl_lahir"`
	Jenis_Kelamin string    `gorm:"type:VARCHAR(100);column:jenis_kelamin" json:"jenis_kelamin"`
	Pekerjaan     string    `gorm:"type:VARCHAR(100);column:pekerjaan" json:"pekerjaan"`
	KD_Pendidikan string    `gorm:"type:VARCHAR(10);column:kd_pendidikan" json:"kd_pendidikan"`
	KD_Agama      string    `gorm:"type:VARCHAR(10);column:kd_agama" json:"kd_agama"`
	KD_Hub        string    `gorm:"type:VARCHAR(10);column:kd_hub_keluarga" json:"kd_hub_keluarga"`
	Info          string    `gorm:"type:VARCHAR(300);column:keterangan" json:"keterangan"`
	Status        string    `gorm:"type:CHAR(1);column:status" json:"status"`
	Serial_Number int       `gorm:"type:INT;column:no_urut" json:"no_urut"`
	Created_At    time.Time `gorm:"type:DATETIME;column:tgl_rekam" json:"tgl_rekam"`
	Creator       string    `gorm:"type:VARCHAR(30);column:petugas_rekam" json:"petugas_rekam"`
	Updated_At    time.Time `gorm:"type:DATETIME;column:tgl_ubah" json:"tgl_ubah"`
	Updater       string    `gorm:"type:VARCHAR(30);column:petugas_ubah" json:"petugas_ubah"`
}

type CreateWargaKKRequest struct {
	KD_Warga      string    `json:"warga_id" binding:"required"`
	NIK           string    `json:"nik" binding:"required"`
	Nama_Warga    string    `json:"nama_warga" binding:"required"`
	Tempat_Lahir  string    `json:"tempat_lahir" binding:"required"`
	Tgl_Lahir     time.Time `json:"tgl_lahir" binding:"required"`
	Jenis_Kelamin string    `json:"jenis_kelamin" binding:"required"`
	Pekerjaan     string    `json:"pekerjaan" binding:"required"`
	KD_Pendidikan string    `json:"kd_pendidikan" binding:"required"`
	KD_Agama      string    `json:"kd_agama" binding:"required"`
	KD_Hub        string    `json:"kd_hub_keluarga" binding:"required"`
	Info          string    `json:"keterangan,omitempty"`
	Status        string    `json:"status,omitempty"`
	Serial_Number int       `json:"no_urut,omitempty"`
}

type UpdateWargaKKRequest struct {
	KD_Warga      string    `json:"warga_id,omitempty"`
	NIK           string    `json:"nik,omitempty"`
	Nama_Warga    string    `json:"nama_warga,omitempty"`
	Tempat_Lahir  string    `json:"tempat_lahir,omitempty"`
	Tgl_Lahir     time.Time `json:"tgl_lahir,omitempty"`
	Jenis_Kelamin string    `json:"jenis_kelamin,omitempty"`
	Pekerjaan     string    `json:"pekerjaan,omitempty"`
	KD_Pendidikan string    `json:"kd_pendidikan,omitempty"`
	KD_Agama      string    `json:"kd_agama,omitempty"`
	KD_Hub        string    `json:"kd_hub_keluarga,omitempty"`
	Info          string    `json:"keterangan,omitempty"`
	Status        string    `json:"status,omitempty"`
	Serial_Number int       `json:"no_urut,omitempty"`
}

func (wk *WargaKK) TableName() string {
	return "tr_warga_kk"
}
