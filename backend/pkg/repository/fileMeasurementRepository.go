package repository

import (
	"encoding/json"
	"os"
	"path"

	"github.com/rs/homecontrol/pkg/models"
)

func SaveLatestAll(folder string, measurements []models.Measurement) error {
	filePath := getLatestFilePath(folder)

	jsonString, err := json.Marshal(measurements)
	if err != nil {
		return err
	}

	data := []byte(jsonString)

	return os.WriteFile(filePath, data, 0644)
}

func ReadLatest(folder string) ([]models.Measurement, error) {
	measurements := []models.Measurement{}

	filePath := getLatestFilePath(folder)

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

func ReadLatestByRoomId(folder string, roomId int) (models.Measurement, error) {
	measurement := models.Measurement{}
	measurements, err := ReadLatest(folder)
	if err != nil {
		return measurement, err
	}

	return measurements[0], nil
}

func getLatestFilePath(folder string) string {
	file := "latest.txt"
	return path.Join(folder, file)
}
