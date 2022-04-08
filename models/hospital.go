package models

import "github.com/jinzhu/gorm"

type Hospital struct {
	gorm.Model
	Patient_name string
	Patient_id   string
	Room_no      int
	Disease      string
}
