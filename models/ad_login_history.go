package models

import "time"

type LoginHistory struct {
	ID          int       `gorm:"type:INT;column:id;autoIncrement;primaryKey;" json:"id"`
	KD_User     string    `gorm:"type:VARCHAR(100);column:kd_user;primaryKey" json:"kd_user"`
	User        Users     `gorm:"foreignKey:KD_User;references:ID_User" json:"user"`
	Login_Time  time.Time `gorm:"type:DATETIME;column:waktu_login" json:"waktu_login"`
	Logout_Time time.Time `gorm:"type:DATETIME;column:waktu_logout" json:"waktu_logout"`
	IP          string    `gorm:"type:VARCHAR(100);column:ip" json:"ip"`
	Browser     string    `gorm:"type:VARCHAR(100);column:browser" json:"browser"`
	Host_Name   string    `gorm:"type:VARCHAR(100);column:host_name" json:"host_name"`
	Info        string    `gorm:"type:VARCHAR(300);column:keterangan" json:"keterangan"`
	Status      string    `gorm:"type:CHAR(1);column:status" json:"status"`
	Created_At  time.Time `gorm:"type:DATETIME;column:tgl_rekam" json:"tgl_rekam"`
	Creator     string    `gorm:"type:VARCHAR(30);column:petugas_rekam" json:"petugas_rekam"`
	Updated_At  time.Time `gorm:"type:DATETIME;column:tgl_ubah" json:"tgl_ubah"`
	Updater     string    `gorm:"type:VARCHAR(30);column:petugas_ubah" json:"petugas_ubah"`
}

func (lh *LoginHistory) TableName() string {
	return "ad_login_history"
}
