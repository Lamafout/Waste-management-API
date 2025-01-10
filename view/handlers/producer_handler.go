package handler

import (
	"encoding/json"
	"net/http"
	controller "waste_management/controller/controller_interfaces"
)

type ProducerHandler struct {
	controller controller.ProducerController
}

func NewProducerHandler(controller controller.ProducerController) *ProducerHandler {
	return &ProducerHandler{controller}
}

func (h *ProducerHandler) GetProducers(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("filter")
	page := r.URL.Query().Get("page")

	producers, err := h.controller.GetProducers(filter, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(producers); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}

func (h *ProducerHandler) PostProducer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var producer map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&producer); err != nil {
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	err := h.controller.PostProducer(producer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}