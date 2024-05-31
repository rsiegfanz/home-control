package httpfetcher

import (
	"log"

	"github.com/rs/homecontrol/pkg/models"
)

func MapMeasurementDtosToModels(dtos []MeasurementDto) []models.Measurement {
	result := []models.Measurement{}
	for _, dto := range dtos {
		obj := MapMeasurementDtoToModel(dto)
		result = append(result, obj)
	}
	return result
}

func MapMeasurementDtoToModel(dto MeasurementDto) models.Measurement {
	id := 0
	switch dto.FileId {
	case "funkbme280": // Garage
		id = 1
		temp := dto.Temperature
		dto.Temperature = dto.Humidity
		dto.Humidity = temp

	case "acbfbd86e3e.werte": // EG Wohnzimmer
		id = 10
	case "8cce4ef1ddea.werte": // EG Küche
		id = 11

	case "4091514f739a.werte": // 1 OG Wohnzimmer
		id = 20
	case "4091514ef9c0.werte": // 1 OG Küche
		id = 21
	case "e89f6d94fb4.werte": // 1 OG Schlafzimmer
		id = 22
	case "acbfbd829b2.werte": // 1 OG Empore
		id = 23

	case "e868e758ea4f.werte": // Keller Gästezimmer
		id = 30
	case "8cce4ef2882a.werte": // Keller Fitnessraum
		id = 31
	default:
		log.Println("Unknown FileId: ", dto.FileId)
	}

	return models.Measurement{Id: id, Temperature: dto.Temperature, Humidity: dto.Humidity}
}
