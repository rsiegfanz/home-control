package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/rsiegfanz/home-control/backend/server/pkg/models"
	"gorm.io/gorm"
)

type Repository struct {
}

func CreateInstance(db *gorm.DB) (*Repository, error) {

	return &Repository{}, nil
}

func (r *Repository) SaveLatestAll(measurements []models.Measurement) error {
	jsonString, err := json.Marshal(measurements)
	if err != nil {
		return err
	}

	data := []byte(jsonString)

	return os.WriteFile("", data, 0644)
}

func (r *Repository) ReadLatest() ([]models.Measurement, error) {
	measurements := []models.Measurement{}

	data, err := os.ReadFile("")
	if err != nil {
		return measurements, err
	}

	err = json.Unmarshal(data, &measurements)
	if err != nil {
		return measurements, err
	}

	return measurements, nil
}

func (r *Repository) ReadLatestByRoomId(roomId int) (models.Measurement, error) {
	measurement := models.Measurement{}
	measurements, err := r.ReadLatest()
	if err != nil {
		return measurement, err
	}

	idx := slices.IndexFunc(measurements, func(measurement models.Measurement) bool { return measurement.Id == roomId })

	if idx == -1 {
		return measurement, fmt.Errorf("measurement not found for room %d", roomId)
	}

	return measurements[idx], nil
}

func (r *Repository) createDataFolder() error {
	dir := filepath.Dir("")
	return os.MkdirAll(dir, os.ModePerm)
}
