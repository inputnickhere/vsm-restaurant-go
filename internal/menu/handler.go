package menu

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

type upsertReq struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	IsActive bool   `json:"is_active"`
}

func (h *Handler) RegisterPublic(r chi.Router) {
	r.Get("/menu", h.getPublicMenu)
}

func (h *Handler) RegisterAdmin(r chi.Router) {
	r.Get("/menu", h.adminListAll)
	r.Post("/menu", h.adminCreate)
	r.Put("/menu/{id}", h.adminUpdate)
	r.Delete("/menu/{id}", h.adminDelete)
}

func (h *Handler) getPublicMenu(w http.ResponseWriter, r *http.Request) {
	items, err := h.svc.ListPublic(r.Context())
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (h *Handler) adminListAll(w http.ResponseWriter, r *http.Request) {
	items, err := h.svc.ListAll(r.Context())
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (h *Handler) adminCreate(w http.ResponseWriter, r *http.Request) {
	var req upsertReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	it, err := h.svc.Create(r.Context(), req.Name, req.Price, req.IsActive)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, it)
}

func (h *Handler) adminUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}

	var req upsertReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	it, err := h.svc.Update(r.Context(), id, req.Name, req.Price, req.IsActive)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, it)
}

func (h *Handler) adminDelete(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "bad id", http.StatusBadRequest)
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		writeServiceError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func parseID(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func writeServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrInvalidName), errors.Is(err, ErrInvalidPrice):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		// (позже добавим 404 по rows affected, сейчас MVP)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
