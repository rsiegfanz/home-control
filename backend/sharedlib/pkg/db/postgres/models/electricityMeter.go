package models

import "gorm.io/gorm"

type ElectricityMeter struct {
	gorm.Model
	Id         uint   `gorm:"primaryKey;autoIncrement"`
	ExternalId string `gorm:"type:varchar(20);unique"`
	Name       string `gorm:"type:varchar(100)"`
}
