package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/homecontrol/pkg/repository"
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
	params := mux.Vars(r)
	roomIdStr := params["roomId"]

	roomId, err := strconv.Atoi(roomIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	repository := repository.GetInstance()

	measurement, err := repository.ReadLatestByRoomId(roomId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	presenter.RespondWithData(w, presenter.MeasurementPresenter{measurement.Temperature, measurement.Humidity})
}
