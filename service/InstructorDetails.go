package service

import (
	"CollegeAdministration/models"
	"log"
)

func (ac *Service) InsertInstructorDet(id *models.InstructorDetails) error {
	err := ac.daos.InsertInstructorDetails(id)
	if err != nil {
		log.Println("Error while inserting details")
		return err
	}
	return err
}
