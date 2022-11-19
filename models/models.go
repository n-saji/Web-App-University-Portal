package models

import (
	"time"

	"github.com/google/uuid"
)

type StudentInfo struct {
	Id              uuid.UUID `gorm:"primary_key;type:uuid;unique"`
	Name            string
	RollNumber      string
	Age             int64
	CourseId        uuid.UUID
	MarksId         uuid.UUID
	ClassesEnrolled CourseInfo   `gorm:"foreignKey:CourseId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StudentMarks    StudentMarks `gorm:"foreignKey:MarksId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CourseInfo struct {
	Id         uuid.UUID `gorm:"primary_key;unique;type:uuid;"`
	CourseName string    `gorm:"unique"`
}

type StudentMarks struct {
	Id         uuid.UUID `gorm:"primary_key;unique;type:uuid;"`
	StudentId  uuid.UUID
	CourseId   uuid.UUID
	CourseName string
	Marks      int64
	Grade      string
}

type InstructorDetails struct {
	Id              uuid.UUID `gorm:"primary_key;unique;type:uuid;"`
	InstructorCode  string
	InstructorName  string
	Department      string
	CourseId        uuid.UUID
	CourseName      string
	ClassesEnrolled CourseInfo `gorm:"foreignKey:CourseId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type InstructorLogin struct {
	Id       uuid.UUID `gorm:"primary_key;unique;type:uuid;"`
	EmailId  string
	Password string
}

type Token_generator struct {
	Token     uuid.UUID `gorm:"primary_key;unique;type:uuid;"`
	ValidFrom time.Time
	ValidTill time.Time
	IsValid   bool
}
type StudentsMarksForCourse struct {
	Course_name     string
	StudentId       []string
	StudentNameMark map[string]int64
	Ranking         map[int64]string
}

type StudentSelectiveData struct {
	Name       string
	Course     string
	RollNumber string
}
type DeleteResponse struct {
	Message string
	Courses []CourseInfo
}
