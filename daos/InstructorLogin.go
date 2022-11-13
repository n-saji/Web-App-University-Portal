package daos

import (
	"CollegeAdministration/models"
	"fmt"

	"github.com/google/uuid"
)

func (ac *AdminstrationCloud) CheckIDPresent(id uuid.UUID) error {

	var id_exits string
	err := ac.dbConn.Select("id").Table("instructor_logins").Where("id = ?", id).Find(&id_exits).Error
	if err != nil {
		return err
	}
	if id_exits != "" {
		return fmt.Errorf("already created ")
	}
	return nil
}
func (ac *AdminstrationCloud) StoreCredentialsForInstructor(il models.InstructorLogin) error {

	err := ac.dbConn.Table("instructor_logins").Create(&il).Error

	if err != nil {
		return err
	}
	return nil
}
