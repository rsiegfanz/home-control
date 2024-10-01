package models

import (
	"time"

	"gorm.io/gorm"
)

type ClimateMeasurement struct {
	gorm.Model
	Id             uint      `gorm:"primaryKey;autoIncrement"`
	RecordedAt     time.Time `gorm:"timestamptz;primaryKey;not null;uniqueIndex:idx_timestamp_room"`
	RoomExternalId string    `gorm:"type:varchar(100);not null;uniqueIndex:idx_timestamp_room"`
	Temperature    float32   `gorm:"type:numeric(5,2)"`
	Humidity       float32   `gorm:"type:numeric(5,2)"`
}
