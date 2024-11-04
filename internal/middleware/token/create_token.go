package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTMaker is
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker is
func NewJWTMaker(secretKey string) *JWTMaker {
	return &JWTMaker{secretKey}
}

// CreateToken method is
func (maker *JWTMaker) CreateToken(id int64, email string, isAdmin bool, duration time.Duration) (string, *UserClaims, error) {
	claims, err := NewUserClaims(id, email, isAdmin, duration)
	if err != nil {
		return "", nil, fmt.Errorf("error getting claims")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(maker.secretKey))
	if err != nil {
		return "", nil, fmt.Errorf("error getting tokenStr: %w", err)
	}

	return tokenStr, claims, nil
}

// VerifyToken is
func (maker *JWTMaker) VerifyToken(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// verify the signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(maker.secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token with claims: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims: %w", err)
	}

	return claims, nil
}
