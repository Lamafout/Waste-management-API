package handler

import (
	"encoding/json"
	"net/http"
	controller "waste_management/controller/controller_interfaces"
	"waste_management/view/responses"
)

type TechnologyHandler struct {
	controller controller.TechnologyController
}

func NewTechnologyHandler(controller controller.TechnologyController) *TechnologyHandler {
	return &TechnologyHandler{controller}
}

func (h *TechnologyHandler) PostTechnology(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var technology map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&technology); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.controller.PostTechnology(technology)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *TechnologyHandler) GetTechnology(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	technology, err := h.controller.GetTechnology(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(technology); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}

func (h *TechnologyHandler) GetTechnologies(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("filter")
	page := r.URL.Query().Get("page")

	technologies, amount, err := h.controller.GetTechnologies(filter, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := responses.NewGetTechnologiesResponse(amount, technologies)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}
