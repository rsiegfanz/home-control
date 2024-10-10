package controllers

import (
	"net/http"

	"github.com/rsiegfanz/home-control/backend/server/pkg/rest/presenter"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
)

func (c *Controller) GetRoomsHandler(w http.ResponseWriter, r *http.Request) {

	var rooms []models.Room
	err := c.DB.Find(&rooms).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	roomsPresenter := make([]presenter.RoomPresenter, len(rooms))

	for idx, room := range rooms {
		roomsPresenter[idx] = presenter.RoomPresenter{Id: room.Id, ExternalId: room.ExternalId, Name: room.Name}
	}

	presenter.RespondWithData(w, roomsPresenter)
}
