package database

import (
	"github.com/Ansh5461/HospitalManagement/models"
	"github.com/jinzhu/gorm"
)

func GetPatient(db *gorm.DB) ([]models.Hospital, error) {
	patient := []models.Hospital{}

	query := db.Select("patient.*")

	err := query.Find(&patient).Error

	if err != nil {
		return nil, err
	}
	return patient, nil
}

func GetPatientRoom(db *gorm.DB, name string) (*[]models.Hospital, error) {
	patient := []models.Hospital{}

	err := db.Select("patient.room_num").Group("patient.patient_name").Where("patient.patient_name = ?", name).First(&patient).Error

	if err != nil {
		return nil, err
	}
	return &patient, nil
}
