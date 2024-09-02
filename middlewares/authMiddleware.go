package middlewares

import (
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt"
    "net/http"
    "strings"
)

var jwtSecret = []byte("YourSecretKey")

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
            c.Abort()
            return
        }

        tokenString := strings.Split(authHeader, " ")[1]

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.NewValidationError("Unexpected signing method", jwt.ValidationErrorSignatureInvalid)
            }
            return jwtSecret, nil
        })

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            c.Set("email", claims["email"])
            c.Set("role", claims["role"])
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Next()
    }
}

func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        role, exists := c.Get("role")
        if !exists || role != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
            c.Abort()
            return
        }
        c.Next()
    }
}
