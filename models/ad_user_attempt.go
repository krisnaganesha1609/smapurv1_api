package models

import "time"

type UserAttempt struct {
	KD_User    string    `gorm:"type:VARCHAR(100);column:kd_user;primaryKey" json:"kd_user"`
	User       Users     `gorm:"foreignKey:KD_User;references:ID_User" json:"user"`
	Attempt    int       `gorm:"type:INT;column:attempt;primaryKey" json:"attempt"`
	Info       string    `gorm:"type:VARCHAR(300);column:keterangan" json:"keterangan"`
	Status     string    `gorm:"type:CHAR(1);column:status" json:"status"`
	Created_At time.Time `gorm:"type:DATETIME;column:tgl_rekam" json:"tgl_rekam"`
	Creator    string    `gorm:"type:VARCHAR(30);column:petugas_rekam" json:"petugas_rekam"`
	Updated_At time.Time `gorm:"type:DATETIME;column:tgl_ubah" json:"tgl_ubah"`
	Updater    string    `gorm:"type:VARCHAR(30);column:petugas_ubah" json:"petugas_ubah"`
}

type EmitUserAttemptResponse struct {
	KD_User    string `json:"kd_user"`
	User       Users  `json:"user"`
	Attempt    int    `json:"attempt"`
	Info       string `json:"keterangan,omitempty"`
	Status     string `json:"status,omitempty"`
	Created_At string `json:"tgl_rekam,omitempty"`
	Updated_At string `json:"tgl_ubah,omitempty"`
}

func (ua *UserAttempt) TableName() string {
	return "ad_user_attempt"
}
