package handler

import (
	"encoding/json"
	"net/http"
	controller "waste_management/controller/controller_interfaces"
)

type FkkoHandler struct {
	controller controller.FkkoController
}

func NewFkkoHandler(controller controller.FkkoController) *FkkoHandler {
	return &FkkoHandler{controller: controller}
}

func (h *FkkoHandler) GetFkkos(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("filter")

	fkkos, err := h.controller.GetFkkos(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(fkkos); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}