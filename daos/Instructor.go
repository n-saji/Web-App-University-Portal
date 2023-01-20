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
		return fmt.Errorf("error while inserting instructor details %s", err.Error())
	}
	return nil
}

func (ac *AdminstrationCloud) GetAllInstructor() ([]*models.InstructorDetails, error) {
	var id []*models.InstructorDetails
	err := ac.dbConn.Order("instructor_name ASC").Find(&id).Error
	if err != nil {
		return nil, fmt.Errorf("not able to retrieve instructor details")
	}
	return id, nil
}
func (ac AdminstrationCloud) GetInstructorDetail(id_exits *models.InstructorDetails) (*models.InstructorDetails, error) {
	var id models.InstructorDetails
	err := ac.dbConn.Where(&id_exits).Find(&id).Error
	if err != nil {
		return nil, fmt.Errorf("record not found")
	}
	return &id, nil
}

func (ac *AdminstrationCloud) DeleteInstructor(name string) error {

	err := ac.dbConn.Where("instructor_name = ?", name).Delete(models.InstructorDetails{}).Error

	if err != nil {
		return err
	}
	return nil
}

func (ac *AdminstrationCloud) GetInstructorWithName(name string) (*models.InstructorDetails, error) {

	var is models.InstructorDetails
	err := ac.dbConn.Model(models.InstructorDetails{}).Select("*").Where("instructor_name = ?", name).Find(&is).Error
	if err != nil {
		return nil, err
	}
	return &is, nil
}

func (ac *AdminstrationCloud) GetInstructorWithSpecifics(condition models.InstructorDetails) ([]*models.InstructorDetails, error) {

	var is []*models.InstructorDetails
	err := ac.dbConn.Model(models.InstructorDetails{}).Select("*").Where(condition).Find(&is).Error
	if err != nil {
		return nil, err
	}
	return is, nil
}

func (ac *AdminstrationCloud) UpdateInstructor(req_id models.InstructorDetails, condition models.InstructorDetails) error {

	log.Println(condition)
	q := ac.dbConn.Model(models.InstructorDetails{}).Where(condition).Updates(req_id)
	if q.Error != nil {
		return q.Error
	}
	return nil
}
