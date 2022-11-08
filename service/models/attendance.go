package models

import (
	"gorm.io/gorm"
	"time"
)

type Attendance struct {
	ID        int            `gorm:"primarykey" json:"id"`
	Status    string         `json:"status"`
	StudentId int            `json:"student_id"`
	Student   User           `gorm:"foreignKey:student_id" json:"student"`
	SubjectID int            `json:"subject_id"`
	Subject   Subject        `gorm:"foreignKey:subject_id" json:"subject,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type AttendancePagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

type AttendancePaginatedData struct {
	Total int `json:"total"`
	data  []User
}
