package handlers

import (
	"encoding/json"
	schemas "image_service/internal/schemas"
	"net/http"
)

type HealtchCheckHandler struct{}

func (hch *HealtchCheckHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encodeError := json.NewEncoder(w).Encode(schemas.NewServerIsAlive())
	if encodeError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewHealtchCheckHandler() *HealtchCheckHandler {
	return &HealtchCheckHandler{}
}
