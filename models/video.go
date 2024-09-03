package models

import (
	"time"
)

type Video struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"not null"`
	Category    string    `gorm:"not null"`
	CreatedBy   string    `gorm:"not null"`
	TimeCreated time.Time `gorm:"not null"`
	Description string    `gorm:"not null"`
	Views       int       `gorm:"default:0"`
	Comments    []Comment // Assuming you have a Comment model
	Subscribers []Subscriber // Assuming you have a Subscriber model
}

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"not null"`
	VideoID uint   `gorm:"not null"`
}

type Subscriber struct {
	ID     uint   `gorm:"primaryKey"`
	UserID string `gorm:"not null"`
	VideoID uint  `gorm:"not null"`
}
