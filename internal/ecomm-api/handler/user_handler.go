package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jumayevgadam/ecomm_mysql/internal/models"
	"github.com/jumayevgadam/ecomm_mysql/pkg/utils"
)

// CreateUser method is
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u UserReq
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "error decoding user", 400)
		return
	}

	// hash password
	hashed, err := utils.HashPassword(u.Password)
	if err != nil {
		http.Error(w, "error hashing password", 400)
		return
	}
	u.Password = hashed

	created, err := h.server.CreateUser(h.ctx, toStorerUser(u))
	if err != nil {
		http.Error(w, "error creating user", 500)
		return
	}

	res := toUserRes(created)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

// LoginUser method is
func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var u LoginUserReq
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "error decoding LoginUserReq", 400)
		return
	}

	gu, err := h.server.GetUser(h.ctx, u.Email)
	if err != nil {
		http.Error(w, "error getting user", 500)
		return
	}

	err = utils.CheckPassword(u.Password, gu.Password)
	if err != nil {
		http.Error(w, "error checking password", http.StatusUnauthorized)
		return
	}

	// create a json web token (JWT) and return it as response
	accessToken, accessClaims, err := h.tokenMaker.CreateToken(gu.ID, gu.Email, gu.IsAdmin, 15*time.Minute)
	if err != nil {
		http.Error(w, "error creating token", http.StatusUnauthorized)
		return
	}

	// create a refresh token
	refreshToken, refreshClaims, err := h.tokenMaker.CreateToken(gu.ID, gu.Email, gu.IsAdmin, 24*time.Hour)
	if err != nil {
		http.Error(w, "errorcreating refresh_token", http.StatusUnauthorized)
		return
	}

	// create session this place
	session, err := h.server.CreateSession(h.ctx, &models.Session{
		ID:           refreshClaims.RegisteredClaims.ID,
		UserEmail:    gu.Email,
		RefreshToken: refreshToken,
		IsRevoked:    false,
		ExpiresAt:    refreshClaims.ExpiresAt.Time,
	})
	if err != nil {
		http.Error(w, "error creating session", http.StatusInternalServerError)
		return
	}

	res := LoginUserRes{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessClaims.ExpiresAt.Time,
		RefreshTokenExpiresAt: refreshClaims.ExpiresAt.Time,
		User:                  toUserRes(gu),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// LogOut method is
func (h *Handler) LogOutUser(w http.ResponseWriter, r *http.Request) {
	// we will later get the session ID from the token payload of the authenticated user
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing session ID", http.StatusBadRequest)
		return
	}

	if err := h.server.DeleteSession(h.ctx, id); err != nil {
		http.Error(w, "error deleting session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RenewAccessToken method is
func (h *Handler) RenewAccessToken(w http.ResponseWriter, r *http.Request) {
	var req RenewAccessTokenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	refreshClaims, err := h.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		http.Error(w, "error verifying token", http.StatusUnauthorized)
		log.Println(err)
		return
	}

	session, err := h.server.GetSession(h.ctx, refreshClaims.RegisteredClaims.ID)
	if err != nil {
		http.Error(w, "error getting session", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if session.IsRevoked {
		http.Error(w, "session revoked", http.StatusUnauthorized)
		return
	}

	if session.UserEmail != refreshClaims.Email {
		http.Error(w, "invalid session", http.StatusUnauthorized)
	}

	accessToken, accessClaims, err := h.tokenMaker.CreateToken(
		refreshClaims.ID,
		refreshClaims.Email,
		refreshClaims.IsAdmin,
		15*time.Minute,
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "error creating token", http.StatusInternalServerError)
	}

	res := RenewAccessTokenRes{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessClaims.ExpiresAt.Time,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// RevokeSession method is
func (h *Handler) RevokeSession(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing session ID", http.StatusBadRequest)
		return
	}

	if err := h.server.RevokeSession(h.ctx, id); err != nil {
		http.Error(w, "error revoking session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ListUsers method is
func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.server.ListUsers(h.ctx)
	if err != nil {
		http.Error(w, "error listing users", 500)
		return
	}

	var res ListUserRes
	for _, u := range users {
		res.Users = append(res.Users, toUserRes(&u))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// UpdateUser method is
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var u UserReq
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "error decoding request body", 400)
		return
	}

	user, err := h.server.GetUser(h.ctx, u.Email)
	if err != nil {
		http.Error(w, "error getting user", 500)
		return
	}

	// patch our user request
	patchUserReq(user, u)

	updated, err := h.server.UpdateUser(h.ctx, user)
	if err != nil {
		http.Error(w, "error updating user", 500)
		return
	}

	res := toUserRes(updated)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// DeleteUser method is
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "error parsing ID", 400)
		return
	}

	err = h.server.DeleteUser(h.ctx, i)
	if err != nil {
		http.Error(w, "error deleting user", 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
