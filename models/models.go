package models

import (
	"database/sql/driver"
	"encoding/json"

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
	CourseName string    `gorm:"unique" json:"course_name"`
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
	Id              uuid.UUID    `gorm:"primary_key;unique;type:uuid" json:"id"`
	InstructorCode  string       `json:"instructor_code"`
	InstructorName  string       `json:"instructor_name"`
	Department      string       `json:"department"`
	CourseId        uuid.UUID    `json:"course_id"`
	CourseName      string       `json:"course_name"`
	ClassesEnrolled CourseInfo   `gorm:"foreignKey:course_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StudentsList    StudentsList `gorm:"type:jsonb;" json:"students_list"`
}

type InstructorLogin struct {
	Id       uuid.UUID `gorm:"primary_key;unique;type:uuid;"`
	EmailId  string
	Password string
}

type Token_generator struct {
	Token     uuid.UUID `gorm:"primary_key;unique;type:uuid;"`
	ValidFrom int64
	ValidTill int64
	IsValid   bool
}

// not part of db
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

type StudentsList struct {
	Info []StudentInfo `json:"info"`
}

func (j StudentsList) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *StudentsList) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}
