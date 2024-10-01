package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rsiegfanz/home-control/backend/server/pkg/rest/presenter"
)

func (c *Controller) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (c *Controller) GetRoomsHandler(w http.ResponseWriter, r *http.Request) {
	rooms := []presenter.RoomPresenter{
		{1, "Büro"},
		{2, "Wohnzimmer"},
		{3, "Küche"},
		{4, "Schlafzimmer"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rooms)
}
