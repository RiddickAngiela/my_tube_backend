package utils

import (
    "github.com/dgrijalva/jwt-go"
    "os"
    "time"
)

// Exported variable for JwtSecret
var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": email,
        "exp":   time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString(JwtSecret)
}
