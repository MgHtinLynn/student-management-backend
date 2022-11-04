package models

import (
	"gorm.io/gorm"
	"time"
)

type Subject struct {
	ID        int            `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	TeacherId int            `json:"teacher_id"`
	Teacher   User           `gorm:"foreignKey:teacher_id" json:"teacher"`
	LectureId int            `json:"lecture_id"`
	Lecture   Lecture        `gorm:"foreignKey:lecture_id" json:"lecture,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type SubjectPagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

type SubjectPaginatedData struct {
	Total int `json:"total"`
	data  []User
}
