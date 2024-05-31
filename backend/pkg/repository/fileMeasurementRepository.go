package repository

import (
	"encoding/json"
	"os"

	"github.com/rs/homecontrol/pkg/models"
)

func SaveLatestAll(filePath string, measurements []models.Measurement) error {
	jsonString, err := json.Marshal(measurements)
	if err != nil {
		return err
	}

	data := []byte(jsonString)

	return os.WriteFile(filePath, data, 0644)
}

func ReadLatest(filePath string) ([]models.Measurement, error) {
	measurements := []models.Measurement{}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return measurements, err
	}

	err = json.Unmarshal(data, &measurements)
	if err != nil {
		return measurements, err
	}

	return measurements, nil
}

func ReadLatestByRoomId(filePath string, roomId int) (models.Measurement, error) {
	measurement := models.Measurement{}
	measurements, err := ReadLatest(filePath)
	if err != nil {
		return measurement, err
	}

	return measurements[0], nil
}
