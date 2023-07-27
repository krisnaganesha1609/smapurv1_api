package models

import "time"

type Warga struct {
	ID_Warga      string      `gorm:"type:VARCHAR(10);column:warga_id;primaryKey" json:"warga_id"`
	KD_Banjar     string      `gorm:"type:VARCHAR(10);column:kd_banjar" json:"kd_banjar"`
	Banjar        Banjar      `gorm:"foreignKey:KD_Banjar;references:ID_Banjar" json:"banjar"`
	KD_Tempek     string      `gorm:"type:VARCHAR(10);column:kd_tempek" json:"kd_tempek"`
	Tempek        Tempek      `gorm:"foreignKey:KD_Tempek;references:ID_Tempek" json:"tempek"`
	Nama_Warga    string      `gorm:"type:VARCHAR(100);column:nama_warga" json:"nama_warga"`
	NIK           string      `gorm:"type:VARCHAR(100);column:nik" json:"nik"`
	Tempat_Lahir  string      `gorm:"type:VARCHAR(100);column:tempat_lahir" json:"tempat_lahir"`
	Tgl_Lahir     time.Time   `gorm:"type:DATETIME;column:tgl_lahir" json:"tgl_lahir"`
	Jenis_Kelamin string      `gorm:"type:VARCHAR(100);column:jenis_kelamin" json:"jenis_kelamin"`
	Pekerjaan     string      `gorm:"type:VARCHAR(100);column:pekerjaan" json:"pekerjaan"`
	KD_Pendidikan string      `gorm:"type:VARCHAR(10);column:kd_pendidikan" json:"kd_pendidikan"`
	Pendidikan    Pendidikan  `gorm:"foreignKey:KD_Pendidikan;references:ID_Pendidikan" json:"pendidikan"`
	Alamat1       string      `gorm:"type:VARCHAR(800);column:alamat1" json:"alamat1"`
	Alamat2       string      `gorm:"type:VARCHAR(800);column:alamat2" json:"alamat2"`
	No_Telp       string      `gorm:"type:VARCHAR(100);column:no_telp" json:"no_telp"`
	No_Hp         string      `gorm:"type:VARCHAR(100);column:no_hp" json:"no_hp"`
	Email         string      `gorm:"type:VARCHAR(100);column:email" json:"email"`
	KD_Hub        string      `gorm:"type:VARCHAR(10);column:kd_hub_keluarga" json:"kd_hub_keluarga"`
	Hubungan      HubKeluarga `gorm:"foreignKey:KD_Hub;references:ID_Hubungan" json:"hub_keluarga"`
	KD_Agama      string      `gorm:"type:VARCHAR(10);column:kd_agama" json:"kd_agama"`
	Agama         Agama       `gorm:"foreignKey:KD_Agama;references:ID_Agama" json:"agama"`
	Info          string      `gorm:"type:VARCHAR(300);column:keterangan" json:"keterangan"`
	Status        string      `gorm:"type:CHAR(1);column:status" json:"status"`
	Serial_Number int         `gorm:"type:INT;column:no_urut" json:"no_urut"`
	Created_At    time.Time   `gorm:"type:DATETIME;column:tgl_rekam" json:"tgl_rekam"`
	Creator       string      `gorm:"type:VARCHAR(30);column:petugas_rekam" json:"petugas_rekam"`
	Updated_At    time.Time   `gorm:"type:DATETIME;column:tgl_ubah" json:"tgl_ubah"`
	Updater       string      `gorm:"type:VARCHAR(30);column:petugas_ubah" json:"petugas_ubah"`
}

type CreateWargaRequest struct {
	ID_Warga      string    `json:"warga_id" binding:"required"`
	KD_Banjar     string    `json:"kd_banjar" binding:"required"`
	KD_Tempek     string    `json:"kd_tempek" binding:"required"`
	Nama_Warga    string    `json:"nama_warga" binding:"required"`
	NIK           string    `json:"nik" binding:"required"`
	Tempat_Lahir  string    `json:"tempat_lahir" binding:"required"`
	Tgl_Lahir     time.Time `json:"tgl_lahir" binding:"required"`
	Jenis_Kelamin string    `json:"jenis_kelamin" binding:"required"`
	Pekerjaan     string    `json:"pekerjaan" binding:"required"`
	KD_Pendidikan string    `json:"kd_pendidikan" binding:"required"`
	Alamat1       string    `json:"alamat1" binding:"required"`
	Alamat2       string    `json:"alamat2" binding:"required"`
	No_Telp       string    `json:"no_telp" binding:"required"`
	No_Hp         string    `json:"no_hp" binding:"required"`
	Email         string    `json:"email" binding:"required"`
	KD_Hub        string    `json:"kd_hub_keluarga" binding:"required"`
	KD_Agama      string    `json:"kd_agama" binding:"required"`
	Info          string    `json:"keterangan,omitempty"`
	Status        string    `json:"status,omitempty"`
	Serial_Number int       `json:"no_urut,omitempty"`
}

type UpdateWargaRequest struct {
	KD_Banjar     string    `json:"kd_banjar,omitempty"`
	KD_Tempek     string    `json:"kd_tempek,omitempty"`
	Nama_Warga    string    `json:"nama_warga,omitempty"`
	NIK           string    `json:"nik,omitempty"`
	Tempat_Lahir  string    `json:"tempat_lahir,omitempty"`
	Tgl_Lahir     time.Time `json:"tgl_lahir,omitempty"`
	Jenis_Kelamin string    `json:"jenis_kelamin,omitempty"`
	Pekerjaan     string    `json:"pekerjaan,omitempty"`
	KD_Pendidikan string    `json:"kd_pendidikan,omitempty"`
	Alamat1       string    `json:"alamat1,omitempty"`
	Alamat2       string    `json:"alamat2,omitempty"`
	No_Telp       string    `json:"no_telp,omitempty"`
	No_Hp         string    `json:"no_hp,omitempty"`
	Email         string    `json:"email,omitempty"`
	KD_Hub        string    `json:"kd_hub_keluarga,omitempty"`
	KD_Agama      string    `json:"kd_agama,omitempty"`
	Info          string    `json:"keterangan,omitempty"`
	Status        string    `json:"status,omitempty"`
	Serial_Number int       `json:"no_urut,omitempty"`
}

func (w *Warga) TableName() string {
	return "tr_warga"
}
