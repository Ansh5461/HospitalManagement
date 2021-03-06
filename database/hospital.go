package database

import (
	"github.com/Ansh5461/HospitalManagement/models"
	"gorm.io/gorm"
)

func GetPatient(db *gorm.DB) ([]models.Hospital, error) {
	patient := []models.Hospital{}

	query := db.Select("hospital.*")

	err := query.Find(&patient).Error

	if err != nil {
		return nil, err
	}
	return patient, nil
}

func GetPatientRoom(db *gorm.DB, name string) (*[]models.Hospital, error) {
	patient := []models.Hospital{}

	err := db.Select("hospital.room_num").Group("hospital.patient_name").Where("hospital.patient_name = ?", name).First(&patient).Error

	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func GetPatientByID(db *gorm.DB, id string) (*models.Hospital, error) {
	p := models.Hospital{}
	err := db.Select("hospital.*").Group("hospital.patient_id").Where("hospital.patient_id = ?", id).First(&p).Error

	if err != nil {
		return nil, err
	}
	return &p, nil
}

func DeletePatientByID(db *gorm.DB, id string) error {
	var p models.Hospital
	err := db.Where("hospital.patient_id = ?", id).Delete(&p).Error
	if err != nil {
		return err
	}
	return nil
}

func SavePatient(db *gorm.DB, p *models.Hospital) error {
	err := db.Save(p).Error
	if err != nil {
		return err
	}
	return nil
}
