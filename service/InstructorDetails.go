package service

import (
	"CollegeAdministration/models"
	"fmt"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (ac *Service) InsertInstructorDetails(iid *models.InstructorDetails) (uuid.UUID, error) {
	cn, err1 := ac.daos.GetCourseByName(iid.CourseName)
	if err1 != nil {
		return uuid.Nil, fmt.Errorf("course not available")
	}
	iid.CourseId = cn.Id
	cd_exist, _ := ac.daos.GetInstructorDetail(iid)
	if cd_exist.Id != uuid.Nil {
		return uuid.Nil, fmt.Errorf("instructor exits")
	}
	ok, err2 := ac.ValidateInstructorDetails(iid)

	if !ok {
		return uuid.Nil, err2
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

func (s *Service) GetInstructorDetailsWithConditions(order_clause string) ([]*models.InstructorDetails, error) {

	id, err := s.daos.GetAllInstructorOrderByCondition(order_clause)
	for _, eachId := range id {
		eachId.ClassesEnrolled, _ = s.daos.GetCourseByName(eachId.CourseName)
	}
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (ac *Service) StoreInstructoLogindetails(id uuid.UUID, emailid, password string) error {

	var credentials models.InstructorLogin

	crypted_password, err2 := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err2 != nil {
		return fmt.Errorf("error parsing password")
	}
	credentials.Id = id
	credentials.EmailId = emailid
	credentials.Password = string(crypted_password)

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

func (s *Service) Update_Instructor(req_id *models.InstructorDetails, cond models.InstructorDetails) error {

	list_details, err := s.GetInstructorDetailWithSpecifics(cond)
	if err != nil {
		return fmt.Errorf(fmt.Sprint("fetching error " + err.Error()))
	}
	if len(list_details) == 0 {
		return fmt.Errorf("no instructor found with given details")
	}
	if req_id.CourseName != "" {
		status := s.daos.CheckCourse(req_id.CourseName)
		if !status {
			return fmt.Errorf("course does not exits")
		}
		course_details, _ := s.daos.GetCourseByName(req_id.CourseName)
		req_id.CourseId = course_details.Id
	}

	ok, err2 := s.ValidateInstructorDetails(req_id)

	if !ok {
		return err2
	}

	err1 := s.daos.UpdateInstructor(req_id, cond)

	if err1 != nil {
		return err1
	}

	return nil

}
func (s *Service) GetInstructorDetailWithSpecifics(req models.InstructorDetails) ([]*models.InstructorDetails, error) {

	id_list, err := s.daos.RetieveInstructorDetailsWithCondition(req)
	if err != nil {
		return nil, fmt.Errorf("error %s", err.Error())
	}

	return id_list, nil

}

func (s *Service) DeleteInstructorWithConditions(id_condition *models.InstructorDetails) error {

	id_list, err := s.daos.RetieveInstructorDetailsWithCondition(*id_condition)
	if err != nil {
		return err
	}
	for _, each_id := range id_list {
		if each_id.Id == uuid.Nil {
			return fmt.Errorf("instructor not found")
		}
		err1 := s.daos.DeleteInstructorLogin(each_id.Id)
		if err1 != nil {
			return err1
		}
		err2 := s.daos.DeleteInstructorWithConditions(each_id)
		if err2 != nil {
			return err2
		}
	}

	return nil
}

func (s *Service) GetInstructorIDWithEmail(email string) (string, error) {

	instructor_id, err := s.daos.GetIDUsingEmail(email)
	if err != nil {
		return "", fmt.Errorf("failed fetching :" + err.Error())
	}
	return instructor_id, nil
}

func (s *Service) GetInstructorNamewithId(id string) (*models.InstructorDetails, error) {
	iid := &models.InstructorDetails{}
	id_uuid, err1 := uuid.Parse(id)
	if err1 != nil {
		return nil, err1
	}
	iid.Id = id_uuid
	i_details, err := s.daos.GetInstructorDetail(iid)
	if err != nil {
		return nil, err
	}
	return i_details, nil
}
