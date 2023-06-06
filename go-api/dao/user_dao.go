package dao

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"column:nombre"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:contrasena"`
	Role     string `gorm:"column:rol"`
}

var (
	// ErrUserNotFound se utiliza cuando no se encuentra un usuario
	ErrUserNotFound = errors.New("user not found")
)
