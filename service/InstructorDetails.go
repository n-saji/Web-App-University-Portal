package service

import (
	"CollegeAdministration/models"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (ac *Service) InsertInstructorDet(iid *models.InstructorDetails) (uuid.UUID, error) {
	cn, err1 := ac.daos.GetCourseByName(iid.CourseName)
	if err1 != nil {
		return uuid.Nil, fmt.Errorf("course not available")
	}
	iid.CourseId = cn.Id
	cd_exist, _ := ac.daos.GetInstructorDetail(iid)
	if cd_exist.Id != uuid.Nil {
		return uuid.Nil, fmt.Errorf("instructor exits")
	}
	iid.Id = uuid.New()
	err := ac.daos.InsertInstructorDetails(iid)
	if err != nil {
		log.Println("Error while inserting details")
		return uuid.Nil, err
	}
	return iid.Id, nil
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

func (ac *Service) StoreInstructoLogindetails(id uuid.UUID, emailid, passowrd string) error {

	var credentials models.InstructorLogin
	credentials.Id = id
	credentials.EmailId = emailid
	credentials.Password = passowrd
	if id == uuid.Nil {
		return fmt.Errorf("uuid cant be null")
	}
	err := ac.daos.CheckIDPresent(credentials.Id)
	if err != nil {
		return err
	}
	err = ac.daos.StoreCredentialsForInstructor(credentials)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteInstructor(name string) error {

	id, err := s.daos.GetInstructorWithName(name)
	if id.Id == uuid.Nil {
		return fmt.Errorf("instructor not found")
	}
	if err != nil {
		return err
	}
	err1 := s.daos.DeleteInstructorLogin(id.Id)
	if err1 != nil {
		return err1
	}

	err2 := s.daos.DeleteInstructor(name)
	if err2 != nil {
		return err2
	}
	return nil
}

func (s *Service) Update_Instructor(req_id models.InstructorDetails, cond models.InstructorDetails)(error){

	err := s.daos.UpdateInstructor(req_id,cond)

	if err != nil {
		return err
	}

	return nil

}