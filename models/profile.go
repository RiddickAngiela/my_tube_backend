// models/profile.go
package models

import "gorm.io/gorm"

type Profile struct {
    gorm.Model
     UserID      uint   `json:"user_id"`
    Name    string `json:"name"`
    Email   string `json:"email"`
    Phone   string `json:"phone"`
    Address string `json:"address"`
    DOB     string `json:"dob"`
}
