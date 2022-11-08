package service

import (
	"CollegeAdministration/models"

	"github.com/google/uuid"
)

func (ac *Service) InsertStudentIdInToMarksTable(cv *models.StudentInfo) (*models.StudentMarks, error) {

	var sm models.StudentMarks

	sm.Id = uuid.New()
	sm.CourseId = cv.ClassesEnrolled.Id
	sm.CourseName = cv.ClassesEnrolled.CourseName
	sm.StudentId = cv.Id
	err := ac.daos.CreateStudentMarks(&sm)
	if err != nil {
		return nil, err
	}
	return &sm, nil
}
