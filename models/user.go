package models

type User struct {
    ID       int    `json:"id"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Role     string `json:"role"` // e.g., "admin" or "user"
}
