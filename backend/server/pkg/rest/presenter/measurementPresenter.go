package presenter

import "time"

type MeasurementPresenter struct {
	RecordedAt  time.Time
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}
