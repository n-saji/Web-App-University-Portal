package service

import (
	"CollegeAdministration/models"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (ac *Service) InsertInstructorDet(iid *models.InstructorDetails) (error, uuid.UUID) {
	cn, err1 := ac.daos.GetCourseByName(iid.CourseName)
	if err1 != nil {
		return fmt.Errorf("course not available"), uuid.Nil
	}
	iid.CourseId = cn.Id
	cd_exist, _ := ac.daos.GetInstructorDetail(iid)
	if cd_exist.Id != uuid.Nil {
		return fmt.Errorf("instructor exits"), uuid.Nil
	}
	iid.Id = uuid.New()
	err := ac.daos.InsertInstructorDetails(iid)
	if err != nil {
		log.Println("Error while inserting details")
		return err, uuid.Nil
	}
	return nil, iid.Id
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
