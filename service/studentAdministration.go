package service

import (
	"CollegeAdministration/models"
	"fmt"

	"github.com/google/uuid"
)

func (ac *Service) InsertValuesToCAd(cv *models.CollegeAdminstration) error {

	cv_id, err := ac.daos.GetCourseByName(cv.ClassesEnrolled.CourseName)
	if err != nil {
		return fmt.Errorf("Course Not Found")
	}
	cv.ClassesEnrolled.Id = cv_id.Id
	ok, err := ac.daos.CheckForRollNo(cv.RollNumber)
	if err != nil {
		return err
	} else {
		if ok {
			return fmt.Errorf("Roll Number already exist!")
		}
	}

	cv.Id = uuid.New()
	err1 := ac.daos.InsertValuesToCollegeAdminstration(cv)
	if err1 != nil {
		return err1
	} else {
		return nil
	}

}

func (ac *Service) RetrieveCAd() ([]*models.CollegeAdminstration, error) {

	rca, err := ac.daos.RetieveCollegeAdminstration()
	if err != nil{
		return rca,err
	}
	return rca, nil
}

func (ac *Service) UpdateCAd(rca *models.CollegeAdminstration) error {
	rc, err1 := ac.daos.GetCourseByName(rca.ClassesEnrolled.CourseName)
	if err1 != nil {
		return fmt.Errorf("Course not Found")
		//rc.CourseName = rca.ClassesEnrolled.CourseName
		//ac.InsertValuesToCA(&rc)
	}

	if rca.Id == uuid.Nil {
		rcaOld, err := ac.daos.GetStudentDetailsByRollNumber(rca.RollNumber)
		if err != nil {
			return fmt.Errorf("ROLL number not found", err)
		}
		rca.Id = rcaOld.Id
	}
	if rca.ClassesEnrolled.Id == uuid.Nil {
		rc, _ = ac.daos.GetCourseByName(rca.ClassesEnrolled.CourseName)
		rca.ClassesEnrolled.Id = rc.Id
		rca.CourseId = rc.Id
	}

	err := ac.daos.UpdateClgStudent(rca)
	if err != nil {
		return err
	}
	return nil

}
