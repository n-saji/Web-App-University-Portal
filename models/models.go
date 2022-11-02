package models

import (
	"github.com/google/uuid"
)

type CollegeAdminstration struct {
	Id              uuid.UUID `gorm:"primary_key;type:uuid;unique"`
	Name            string
	RollNumber      string `gorm:"unique"`
	Age             int64
	CourseId        uuid.UUID
	ClassesEnrolled CoursesAvailable `gorm:"foreignKey:CourseId"`
}

type CoursesAvailable struct {
	Id         uuid.UUID `gorm:"primary_key;unique;type:uuid;"`
	CourseName string    `gorm:"unique"`
}
