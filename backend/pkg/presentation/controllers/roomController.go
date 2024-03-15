package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rs/homecontrol/pkg/presentation/presenter"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func GetRoomsHandler(w http.ResponseWriter, r *http.Request) {
	rooms := []presenter.RoomPresenter{
		{"Büro"},
		{"Wohnzimmer"},
		{"Küche"},
		{"Schlafzimmer"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rooms)
}
