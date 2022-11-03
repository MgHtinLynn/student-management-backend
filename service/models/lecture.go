package models

import (
	"gorm.io/gorm"
	"time"
)

type Lecture struct {
	ID        int            `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	TutorId   int            `json:"tutor_id"`
	Tutor     User           `gorm:"foreignKey:tutor_id" json:"tutor"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type LecturePagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

type LecturePaginatedData struct {
	Total int `json:"total"`
	data  []User
}
