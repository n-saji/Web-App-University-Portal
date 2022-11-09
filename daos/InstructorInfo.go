package daos

import (
	"CollegeAdministration/models"
	"fmt"
	"log"
)

func (ac *AdminstrationCloud) InsertInstructorDetails(id *models.InstructorDetails) error {
	err := ac.dbConn.Table("instructor_details").Create(id).Error
	if err != nil {
		log.Println("Not able to insert instructor details", err)
		return fmt.Errorf("error while inserting instructor details", err)
	}
	log.Println("Stored to database")
	return nil
}
