package session

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Session struct {
	TokenString    string
	ExpirationTime time.Time
}

func GenerateToken(username string) (Session, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return Session{}, err
	}

	return Session{
		TokenString:    tokenString,
		ExpirationTime: expirationTime,
	}, nil
}

func VerifyToken(token string) error {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return err
	}

	if !tkn.Valid {
		return errors.New("Token has expired")
	}

	return nil
}

func GetToken(r *http.Request) (string, error) {
	c, err := r.Cookie("token")

	if err != nil {
		return "", err
	}

	return c.Value, err
}
