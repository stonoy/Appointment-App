package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stonoy/Appointment-App/internal/database"
)

type MyCustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(user database.User, secret string) (string, error) {
	cliams := MyCustomClaims{
		string(user.Role),
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "MyAppointment1",
			Subject:   "the user",
			ID:        fmt.Sprintf("%v", user.ID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	ss, err := token.SignedString([]byte(secret))
	return ss, err
}

func CheckAndExtractToken(token, secret string) (string, string, error) {
	tokenInterface, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", "", err
	}

	if claims, ok := tokenInterface.Claims.(*MyCustomClaims); ok && tokenInterface.Valid {
		role := claims.Role
		id := claims.ID
		return role, id, nil
	} else {
		return "", "", fmt.Errorf("token is not valid")
	}
}
