package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/rs/homecontrol/pkg/config"
	"github.com/rs/homecontrol/pkg/models"
)

var singleInstance *Repository

type Repository struct {
	config *config.Config
}

func CreateInstance(config *config.Config) (*Repository, error) {
	singleInstance = &Repository{config: config}

	err := singleInstance.createDataFolder()
	if err != nil {
		return nil, err
	}

	return singleInstance, nil
}

func GetInstance() *Repository {
	return singleInstance
}

func (r *Repository) SaveLatestAll(measurements []models.Measurement) error {
	jsonString, err := json.Marshal(measurements)
	if err != nil {
		return err
	}

	data := []byte(jsonString)

	return os.WriteFile(r.config.DataPaths.LatestMeasurements, data, 0644)
}

func (r *Repository) ReadLatest() ([]models.Measurement, error) {
	measurements := []models.Measurement{}

	data, err := os.ReadFile(r.config.DataPaths.LatestMeasurements)
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
	dir := filepath.Dir(r.config.DataPaths.LatestMeasurements)
	return os.MkdirAll(dir, os.ModePerm)
}
