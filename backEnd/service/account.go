package service

import (
	"CollegeAdministration/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateNewAccount(acc *models.Account) error {

	err := s.ValidateLogin(acc.Info.Credentials.EmailId, acc.Info.Credentials.Password)
	if err != nil {
		log.Println("failed to validate credentials")
		return err
	}
	err = s.CheckEmailExist(acc.Info.Credentials.EmailId)
	if err != nil {
		log.Println("duplicate email id")
		return err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(acc.Info.Credentials.Password), 10)
	if err != nil {
		log.Println("failed to encrypt password")
		return err
	}
	acc.Info.Credentials.Password = string(password)
	err = s.daos.CreateAccount(acc)
	if err != nil {
		log.Println("failed to create account")
		return err
	}
	allCourses, err := s.RetrieveCA()
	if err != nil || len(allCourses) == 0 {
		log.Println("unable to get any courses")
		return err
	}
	instructorDetails := &models.InstructorDetails{
		Id:              acc.Id,
		InstructorCode:  "-",
		InstructorName:  acc.Name,
		Department:      "Empty Department",
		// CourseId:        allCourses[0].Id,
		// CourseName:      "Empty Course",
		// ClassesEnrolled: models.CourseInfo{},
		// Info:            models.Instructor_Info{},
	}
	err = s.daos.InsertInstructorDetails(instructorDetails)
	if err != nil {
		log.Println("failed to insert into instructor details")
		if err != nil {
			err := s.daos.DeleteAccount(acc.Id)
			if err != nil {
				log.Println("failed to revert account creation changes - delete account")
				return err
			}
		}
		return err
	}
	instructorLogin := models.InstructorLogin{
		Id:       acc.Id,
		EmailId:  acc.Info.Credentials.EmailId,
		Password: acc.Info.Credentials.Password,
	}
	err = s.daos.CreateInstructorLogin(instructorLogin)
	if err != nil {
		log.Println("failed to insert into instructor login")
		if err != nil {
			err := s.daos.DeleteAccount(acc.Id)
			if err != nil {
				log.Println("failed to revert account creation changes - delete account")
				return err
			}
			return err
		}
	}
	return nil
}
