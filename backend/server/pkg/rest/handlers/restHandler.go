package handlers

import "gorm.io/gorm"

type RestHandler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *RestHandler {
	return &RestHandler{DB: db}
}
