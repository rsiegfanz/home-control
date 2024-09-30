package models

import (
	"time"

	"gorm.io/gorm"
)

type ClimateMeasurement struct {
	gorm.Model
	Id             uint      `gorm:"primaryKey;autoIncrement"`
	RoomExternalId string    `gorm:"type:varchar(100);uniqueIndex:idx_room_timestamp"`
	RecordedAt     time.Time `gorm:"timestamptz;uniqueIndex:idx_room_timestamp"`
	Temperature    float32   `gorm:"type:numeric(5,2)"`
	Humidity       float32   `gorm:"type:numeric(5,2)"`
}
