package service

import (
	"CollegeAdministration/models"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (ac *Service) InsertInstructorDet(iid *models.InstructorDetails) error {
	iid.Id = uuid.New()
	cn, err1 := ac.daos.GetCourseByName(iid.CourseName)
	if err1 != nil {
		return fmt.Errorf("course not available")
	}
	iid.CourseId = cn.Id
	err := ac.daos.InsertInstructorDetails(iid)
	if err != nil {
		log.Println("Error while inserting details")
		return err
	}
	return err
}

func (ac *Service) GetInstructorDetails() ([]*models.InstructorDetails, error) {

	id, err := ac.daos.GetAllInstructor()
	for _, eachId := range id {
		eachId.ClassesEnrolled, _ = ac.daos.GetCourseByName(eachId.CourseName)
	}
	if err != nil {
		return nil, err
	}
	return id, nil
}
