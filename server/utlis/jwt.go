package utils

import (
	"time"
	"os"
	"github.com/golang-jwt/jwt/v5"
	"fmt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

var accessSecret  = []byte(os.Getenv("ACCESS_SECRET"))
var refreshSecret = []byte(os.Getenv("REFRESH_SECRET"))

func generateToken(userID int, tokenType string, exp time.Duration, secret []byte) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"type": tokenType,
		"exp":  time.Now().Add(exp).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func GenerateTokens(userID int) (access, refresh string, err error) {
	access, err = generateToken(userID, "access", 15*time.Minute, accessSecret)
	if err != nil {
		return
	}

	refresh, err = generateToken(userID, "refresh", 7*24*time.Hour, refreshSecret)
	return
}

type TokenPayload struct {
	UserID int
	Type   string
}

func ValidateToken(tokenStr string, secret []byte, expectedType string) (*TokenPayload, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	// check exp
	expireTime, err := claims.GetExpirationTime()
	if err == nil || time.Now().After(expireTime.Time) {
		return nil, fmt.Errorf("token expired")
	}
	// check type
	if claims["type"] != expectedType {
		return nil, fmt.Errorf("wrong token type")
	}

	userID := int(claims["sub"].(float64))

	return &TokenPayload{
		UserID: userID,
		Type:   claims["type"].(string),
	}, nil
}