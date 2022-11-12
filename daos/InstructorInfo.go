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

func (ac *AdminstrationCloud) GetAllInstructor() ([]*models.InstructorDetails, error) {
	var id []*models.InstructorDetails
	err := ac.dbConn.Find(&id).Error
	if err != nil {
		return nil, fmt.Errorf("not able to retrieve instructor details")
	}
	return id, nil
}
func (ac AdminstrationCloud) GetInstructorDetail(id_exits *models.InstructorDetails) (*models.InstructorDetails, error) {
	var id models.InstructorDetails
	err := ac.dbConn.Where(&id_exits).Find(&id).Error
	if err != nil {
		fmt.Errorf("record not found")
	}
	return &id, nil
}
