package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

// GenerateJWT generates a JWT token for the given email and role.
func GenerateJWT(email, role string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-default-secret" // Provide a default secret or handle it appropriately
	}

	tokenClaims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT parses the JWT token and returns the claims.
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-default-secret" // Provide a default secret or handle it appropriately
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("Unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, jwt.NewValidationError("Token is expired", jwt.ValidationErrorExpired)
			}
			if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, jwt.NewValidationError("Token not valid yet", jwt.ValidationErrorNotValidYet)
			}
			// Handle other validation errors
			return nil, ve
		}
		// Handle other errors
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.NewValidationError("Invalid token", jwt.ValidationErrorClaimsInvalid)
}
