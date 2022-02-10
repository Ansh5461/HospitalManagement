package database

import (
	"log"

	"github.com/Ansh5461/HospitalManagement/tree/master/models"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "Hospital"
	username := "postgres"
	password := "postgres"
	args := "host=" + host + " port=" + port + " dbName=" + dbName + "username=" + username + " sslmode=disable password=" + password
	db, err := gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Hospital{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
