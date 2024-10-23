package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// CreateOrder is
func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var o OrderReq
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		log.Println(err)
		http.Error(w, "bad request for decoding OrderReq", 400)
		return
	}

	created, err := h.server.CreateOrder(h.ctx, toStorerOrder(o))
	if err != nil {
		http.Error(w, "error create order", 500)
		return
	}

	res := toOrderRes(created)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

// GetOrder is
func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "error parsing ID", 400)
		return
	}

	order, err := h.server.GetOrder(h.ctx, i)
	if err != nil {
		http.Error(w, "error get order", 500)
		return
	}

	res := toOrderRes(order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// ListOrders is
func (h *Handler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.server.ListOrders(h.ctx)
	if err != nil {
		http.Error(w, "error listing products", 500)
		return
	}

	var res []OrderRes
	for _, o := range orders {
		res = append(res, toOrderRes(&o))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// DeleteOrder is
func (h *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "error parsing ID", 400)
		return
	}

	if err := h.server.DeleteOrder(h.ctx, i); err != nil {
		http.Error(w, "error deleting order", 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
