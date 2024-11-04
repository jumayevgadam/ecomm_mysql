package handler

import "time"

// UserReq model is
type UserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

// UserRes model is
type UserRes struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

// ListUserRes model is
type ListUserRes struct {
	Users []UserRes `json:"users"`
}

// LoginUserReq model is
type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginUserRes model is
type LoginUserRes struct {
	SessionID             string    `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	User                  UserRes   `json:"user"`
}

// RenewAccessTokenReq model is
type RenewAccessTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

// RenewAccessTokenRes model is
type RenewAccessTokenRes struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}
