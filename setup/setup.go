package setup

import (
	"fmt"
	"log"
	m "smapurv1_api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(config *Config) {
	var err error
	dbName := config.DBName     //rename the database to match in local mysql "smapur_wsa"
	dbHost := config.DBHost     //rename the DBMS Url to match with your URL "127.0.0.1	"
	dbPort := config.DBPort     //rename the DBMS Port to match with yours "3306"
	dbUser := config.DBUserName //fill the DBMS User "root"
	dbPass := config.DBPassword //fill the DBMS Password ""
	database, err := gorm.Open(mysql.Open(dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Can't Connect To Database!")
	}

	log.Println("Connected To Database!")

	DB = database
}

func Migrations() {

	if err := DB.AutoMigrate(
		&m.LoginHistory{},
		&m.Menu{},
		&m.RoleMenu{},
		&m.Role{},
		&m.Session{},
		&m.UserAttempt{},
		&m.UserRole{},
		&m.Users{},
		&m.Agama{},
		&m.Banjar{},
		&m.HubKeluarga{},
		&m.Pendidikan{},
		&m.SandiTransaksi{},
		&m.Tempek{},
		&m.WargaKK{},
		&m.WargaTransaksi{},
		&m.Warga{},
	); err != nil {
		log.Println(err.Error())
	}

	// Add another triggers or Procedures Here

	fmt.Println("Database Migrated!")
}
