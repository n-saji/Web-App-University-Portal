package service

import (
	"CollegeAdministration/models"
	"fmt"

	"github.com/google/uuid"
)

func (ac *Service) InsertValuesToCA(cv *models.CoursesAvailable) error {

	cv.Id = uuid.New()

	ok := ac.daos.CheckCourse(cv.CourseName)
	if ok {
		return fmt.Errorf("Course Name exits!")
	}
	status := ac.daos.InsertValuesToCoursesAvailable(cv)
	if status != nil {
		return status
	}
	return status

}

func (ac *Service) RetrieveCA() ([]*models.CoursesAvailable, error) {

	rca, err := ac.daos.RetieveCoursesAvailable()
	return rca, err
}


func (ac *Service) UpdateCA(name string, rc *models.CoursesAvailable) error {

	err := ac.daos.UpdateCourseByName(name, rc)
	if err != nil {
		return fmt.Errorf("Not able to update", err)
	}
	return nil
}

