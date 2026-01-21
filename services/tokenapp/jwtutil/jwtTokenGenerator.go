package jwtutil

import (
	"os"
	"time"
	"tokenapp/logger"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(username string, role string) (string, error) {
	logger := logger.LoggingInit()
	logger.Info("Generating token for User: ", username, "having Role: ", role)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "nwapp",
		"aud": role,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	logger.Info("Generating JWT token.")
	SECRET := os.Getenv("SECRET")

	tokenString, err := claims.SignedString([]byte(SECRET))

	if err != nil {
		logger.Error("Failed to generate JWT token.")
		return "", err
	}
	logger.Info("Sucessfully generated JWT token.")
	return tokenString, nil
}

func IsAuthenticated(tokenString string) bool {
	logger := logger.LoggingInit()
	logger.Info("Verifying TOKEN-STRING: ", tokenString)
	SECRET := os.Getenv("SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return false
	}

	if token.Valid {
		logger.Info("TOKEN verification status: ", token.Valid)
		return true
	} else {
		logger.Info("TOKEN verification status: ", token.Valid)
		return false
	}
}

func IsAuthorized(tokenString string, role string) bool {
	logger := logger.LoggingInit()
	logger.Info("Verifying user authorization status for TOKEN: ", tokenString)
	SECRET := os.Getenv("SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		logger.Info("TOKEN is valid.Current role of the user is:", claims["aud"])
		return claims["aud"] == role
	}

	return false
}
