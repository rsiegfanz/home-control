package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/homecontrol/pkg/rest/presenter"
)

func GetTemperaturesByRoomIdHandler(w http.ResponseWriter, r *http.Request) {
	rooms := []presenter.TemperaturePresenter{
		{11.5},
		{11.7},
		{11.9},
		{12.7},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rooms)
}

func GetLatestTemperatureByRoomIdHandler(w http.ResponseWriter, r *http.Request) {
	temperature := presenter.TemperaturePresenter{11.5}

	presenter.RespondWithData(w, temperature)
}
