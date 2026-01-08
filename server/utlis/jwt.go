package utils

import (
	"time"
	"os"
	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"

)

var accessSecretKey = []byte(os.Getenv("ACCESS_SECRET_KEY"))
var refreshSecretKey = []byte(os.Getenv("REFRESH_SECRET_KEY"))


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
=======
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`

	jwt.RegisteredClaims
}

func GenerateTokens(email, userID, userType string) (string, string, error) {
    tokenExpiry := time.Now().Add(24 * time.Hour).Unix()
    refreshTokenExpiry := time.Now().Add(7 * 24 * time.Hour).Unix()
	
    claims := &Claims{
        Email:  email,
        UserID: userID,
        Role:   userType,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Unix(tokenExpiry, 0)),
        },
    }

    refreshClaims := &Claims{
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Unix(refreshTokenExpiry, 0)),
        },
    }

    accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedAccessToken, err := accessTokenObj.SignedString(accessSecretKey)
    if err != nil {
        return "", "", err
    }

    
    refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
    signedRefreshToken, err := refreshTokenObj.SignedString(refreshSecretKey)
    if err != nil {
        return "", "", err
    }

    return signedAccessToken, signedRefreshToken, nil
}

func ValidateAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return accessSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func ValidateRefreshToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return refreshSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func HashPassword(password *string) *string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	hashedPwd := string(bytes)
	return &hashedPwd
} 

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

