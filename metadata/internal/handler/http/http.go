package http

import (
	"encoding/json"
	"errors"
	"log"
	"movieapp/metadata/internal/controller/metadata"
	"movieapp/metadata/internal/repository"
	"net/http"
)

type Handler struct {
	ctrl *metadata.Controller
}

func New(ctrl *metadata.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) GetMetadata(w http.ResponseWriter, r *http.Request) {
	var id = r.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	m, err := h.ctrl.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository got the error %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
