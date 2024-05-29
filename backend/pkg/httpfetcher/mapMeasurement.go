package httpfetcher

import "github.com/rs/homecontrol/pkg/models"

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
	case "4091514f739a.werte":
		id = 1
	case "acbfbd86e3e.werte":
		id = 2
	case "4091514ef9c0.werte":
		id = 3
	case "acbfbd829b2.werte":
		id = 4
	case "8cce4ef1ddea.werte":
		id = 5
	case "e868e758ea4f.werte":
		id = 6
	case "8cce4ef2882a.werte":
		id = 7
	case "e89f6d94fb4.werte":
		id = 8
	case "funkbme280":
		id = 9
		temp := dto.Temperature
		dto.Temperature = dto.Humidity
		dto.Humidity = temp
	}

	return models.Measurement{Id: id, Temperature: dto.Temperature, Humidity: dto.Humidity}
}
