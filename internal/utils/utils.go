package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/gomail.v2"
)

var jwtKey = []byte("yourjwtsecret")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "youremail@example.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.example.com", 587, "youremail@example.com", "yourpassword")

	return d.DialAndSend(m)
}
