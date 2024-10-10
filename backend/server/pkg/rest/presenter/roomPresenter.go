package presenter

type RoomPresenter struct {
	Id         uint   `json:"id"`
	ExternalId string `json:"externalId"`
	Name       string `json:"name"`
}
