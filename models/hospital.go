package models

import "github.com/jinzhu/gorm"

type Hospital struct {
	gorm.Model
	patient_name string
	patient_id   string
	room_no      int
	disease      string
}
