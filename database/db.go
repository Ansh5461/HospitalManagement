package database

import (
	"log"

	"github.com/Ansh5461/HospitalManagement/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "hospital"
	username := "postgres"
	password := "postgres"
	args := "host=" + host + " port=" + port + " user=" + username + " dbname=" + dbName + " sslmode=disable password=" + password
	db, err := gorm.Open(postgres.Open(args), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Hospital{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
