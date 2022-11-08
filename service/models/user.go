package models

import (
	"gorm.io/gorm"
)
import "time"

type TypeTime *time.Time

type User struct {
	ID              int            `gorm:"primarykey" json:"id"`
	Name            string         `json:"name"`
	Email           string         `json:"email"`
	Password        string         `json:"-"`
	Active          bool           `gorm:"type:bool;default:false" json:"active"`
	Role            string         `json:"role"`
	Phone           string         `json:"phone"`
	ProfileUrl      string         `json:"profile_url"`
	Address         string         `json:"address"`
	EmailVerifiedAt TypeTime       `json:"email_verified_at" gorm:"default:null"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

type PaginatedData struct {
	Total int `json:"total"`
	data  []User
}
