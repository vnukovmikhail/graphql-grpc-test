package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Todos []Todo // Один ко многим
}

type Todo struct {
	gorm.Model
	Text   string `gorm:"not null"`
	Done   bool   `gorm:"default:false"`
	UserID uint   // Внешний ключ для GORM (числовой ID)
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Связь с юзером
}
