package models

import (
	"gorm.io/gorm"
	"time"
)

type Certificate struct {
	ID        int            `gorm:"primarykey" json:"id"`
	FilePath  string         `json:"file_path,omitempty" gorm:"default:null"`
	StudentId int            `json:"student_id"`
	Student   User           `gorm:"foreignKey:student_id" json:"student"`
	LectureID int            `json:"lecture_id"`
	Lecture   Lecture        `gorm:"foreignKey:lecture_id" json:"lecture,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type CertificatePagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

type CertificatePaginatedData struct {
	Total int `json:"total"`
	data  []User
}
