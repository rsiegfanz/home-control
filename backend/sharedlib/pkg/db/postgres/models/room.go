package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Id         uint   `gorm:"primaryKey;autoIncrement"`
	ExternalId string `gorm:"type:varchar(100);unique"`
	Name       string `gorm:"type:varchar(100)"`
}
