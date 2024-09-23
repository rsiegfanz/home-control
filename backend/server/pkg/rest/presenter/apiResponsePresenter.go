package presenter

import (
	"encoding/json"
	"net/http"
)

type ApiResponsePresenter struct {
	Status        int         `json:"status"`
	StatusDetails *string     `json:"statusDetails,omitempty"`
	Messages      *[]string   `json:"messages,omitempty"`
	Data          interface{} `json:"data,omitempty"`
}

func NewApiResponsePresenter(status int, statusDetails *string, messages *[]string, data interface{}) *ApiResponsePresenter {
	return &ApiResponsePresenter{
		Status:        status,
		StatusDetails: statusDetails,
		Messages:      messages,
		Data:          data,
	}
}

func CreateWithData(status int, data interface{}, messages *[]string) *ApiResponsePresenter {
	return NewApiResponsePresenter(status, nil, messages, data)
}

func CreateException(status int, messages *[]string, statusDetails *string) *ApiResponsePresenter {
	return NewApiResponsePresenter(status, statusDetails, messages, nil)
}

func RespondWithData(w http.ResponseWriter, data interface{}) {
	apiResponse := CreateWithData(http.StatusOK, data, nil)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResponse)
}
