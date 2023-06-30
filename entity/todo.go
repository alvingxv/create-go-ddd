package entity

import (
	"time"
)

type Todo struct {
	Id        int    `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"not null;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
