package utils

import (
	"errors"
	"time"

	"gofinance/api/pkg/types"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var jwtKey = []byte("your_secret_key_here")

type Claims struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	GoogleID  string    `json:"google_id,omitempty"`
	ExpiredAt time.Time `json:"expired_at"`
	jwt.StandardClaims
}

func GenerateToken(userDTO *types.UserDTO) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		ID:        userDTO.ID,
		Name:      userDTO.Name,
		Email:     userDTO.Email,
		Password:  userDTO.Password,
		GoogleID:  userDTO.GoogleID,
		ExpiredAt: expirationTime,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("token claims are not of type *Claims")
	}

	return claims, nil
}
