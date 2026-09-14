// Package model of user
package model

import "time"

type User struct {
	ID           int64  `gorm:"primaryKey"`
	Email        string `gorm:"uniqueIndex;not null"`
	Username     string `gorm:"not null"`
	PasswordHash string `gorm:"not null"`
	CreatedAt    time.Time
}
