package session

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

const JWT_KEY = "JWT_KEY"

var jwtKey []byte

type Claims struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}

type Session struct {
	TokenString    string
	ExpirationTime time.Time
}

func init() {
	godotenv.Load("../.env")
	jwtKey = []byte(os.Getenv(JWT_KEY))
}

func GenerateToken(id uint) (Session, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Id: id,
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

func VerifyToken(token string) (uint, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return 0, err
	}

	if !tkn.Valid {
		return 0, errors.New("Token has expired")
	}

	return claims.Id, nil
}

func GetToken(r *http.Request) (string, error) {
	c, err := r.Cookie("token")

	if err != nil {
		return "", err
	}

	return c.Value, err
}
