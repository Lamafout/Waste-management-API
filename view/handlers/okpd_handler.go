package handler

import (
	"encoding/json"
	"net/http"
	controller "waste_management/controller/controller_interfaces"
)

type OkpdHandler struct {
	controller controller.OkpdController
}

func NewOkpdHandler(controller controller.OkpdController) *OkpdHandler {
	return &OkpdHandler{controller: controller}
}

func (h *OkpdHandler) GetOkpds(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("filter")

	okpds, err := h.controller.GetOkpds(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(okpds); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}