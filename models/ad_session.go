package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Session_ID uuid.UUID `gorm:"type:VARCHAR(100);column:session_id;primaryKey" json:"session_id"`
	KD_User    string    `gorm:"type:VARCHAR(100);column:kd_user;primaryKey" json:"kd_user"`
	User       Users     `gorm:"foreignKey:KD_User;references:ID_User" json:"user"`
	Expired    time.Time `gorm:"type:DATETIME;column:expired" json:"expired"`
	IP         string    `gorm:"type:VARCHAR(100);column:ip" json:"ip"`
	Info       string    `gorm:"type:VARCHAR(300);column:keterangan" json:"keterangan"`
	Status     string    `gorm:"type:CHAR(1);column:status" json:"status"`
	Created_At time.Time `gorm:"type:DATETIME;column:tgl_rekam" json:"tgl_rekam"`
	Creator    string    `gorm:"type:VARCHAR(30);column:petugas_rekam" json:"petugas_rekam"`
	Updated_At time.Time `gorm:"type:DATETIME;column:tgl_ubah" json:"tgl_ubah"`
	Updater    string    `gorm:"type:VARCHAR(30);column:petugas_ubah" json:"petugas_ubah"`
}

type GetSessionAfterLoginResponse struct {
	Session_ID uuid.UUID `json:"session_id,omitempty"`
	KD_User    string    `json:"kd_user,omitempty"`
	Expired    time.Time `json:"expired,omitempty"`
}

func (se *Session) TableName() string {
	return "ad_session"
}
