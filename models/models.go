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
	Id              uuid.UUID       `gorm:"primary_key;unique;type:uuid" json:"id"`
	InstructorCode  string          `json:"instructor_code"`
	InstructorName  string          `json:"instructor_name"`
	Department      string          `json:"department"`
	CourseId        uuid.UUID       `json:"course_id"`
	CourseName      string          `json:"course_name"`
	ClassesEnrolled CourseInfo      `gorm:"foreignKey:course_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Info            Instructor_Info `gorm:"type:jsonb;" json:"info"`
}

type InstructorLogin struct {
	Id       uuid.UUID `gorm:"primary_key;unique;type:uuid" json:"id"`
	EmailId  string    `json:"email_id"`
	Password string    `json:"password"`
}

type Token_generator struct {
	Token     uuid.UUID `gorm:"primary_key;unique;type:uuid;"`
	ValidFrom int64
	ValidTill int64
	IsValid   bool
}

type Account struct {
	Id   uuid.UUID    `gorm:"primary_key;unique;type=uuid"`
	Name string       `json:"name"`
	Info Account_Info `gorm:"type:jsonb;" json:"info"`
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

type Instructor_Info struct {
	StudentsList []StudentInfo `json:"students_list"`
}

func (j Instructor_Info) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *Instructor_Info) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

type Account_Info struct {
	Credentials InstructorLogin `json:"credentials"`
}

type InstructorProfile struct {
	Name        string
	CourseList  string //needs to be array of string in future
	Department  string
	Code        string
	Credentials InstructorLogin
}
