package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rsiegfanz/home-control/backend/server/pkg/rest/presenter"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
)

func (c *Controller) GetMeasurementsByRoomIdHandler(w http.ResponseWriter, r *http.Request) {
	rooms := []presenter.MeasurementPresenter{
		//		{11.5, 50.0},
		//		{11.7, 49.2},
		//		{11.9, 45.2},
		//		{12.7, 65.0},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rooms)
}

func (c *Controller) GetLatestMeasurementByRoomIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomExternalId := params["roomId"]

	var latestMeasurement models.ClimateMeasurement
	err := c.DB.Where(&models.ClimateMeasurement{RoomExternalId: roomExternalId}).First(&latestMeasurement).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	presenter.RespondWithData(w, presenter.MeasurementPresenter{RecordedAt: latestMeasurement.RecordedAt, Temperature: latestMeasurement.Temperature, Humidity: latestMeasurement.Humidity})
}
