package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/homecontrol/pkg/rest/presenter"
)

func GetMeasurementsByRoomIdHandler(w http.ResponseWriter, r *http.Request) {
	rooms := []presenter.MeasurementPresenter{
		{11.5, 50.0},
		{11.7, 49.2},
		{11.9, 45.2},
		{12.7, 65.0},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rooms)
}

func GetLatestMeasurementByRoomIdHandler(w http.ResponseWriter, r *http.Request) {
	measurement := presenter.MeasurementPresenter{11.5, 34.2}

	presenter.RespondWithData(w, measurement)
}
