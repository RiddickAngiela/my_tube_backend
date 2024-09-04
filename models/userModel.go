package models

import (
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    // Define your user model fields
    Email    string `gorm:"unique"`
    Password string
    Role     string
}

// SetPassword hashes the user's password
func (u *User) SetPassword(password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)
    return nil
}

// CheckPassword compares a hashed password with the provided one
func (u *User) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
