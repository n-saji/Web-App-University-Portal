package daos

import (
	"CollegeAdministration/models"

	"github.com/google/uuid"
)

func (ac *AdminstrationCloud) CreateStudentMarks(sm *models.StudentMarks) error {

	err := ac.dbConn.Table("student_marks").Create(sm).Error

	if err != nil {
		return err
	}
	return nil
}

func (ac *AdminstrationCloud) GetMarksByMarksId(id uuid.UUID) (*models.StudentMarks, error) {

	var sm *models.StudentMarks
	err := ac.dbConn.Model(sm).Where("id = ?", id).Find(&sm).Error

	if err != nil {
		return nil, err
	}
	return sm, nil
}
func (ac *AdminstrationCloud) GetMarksByStudentId(id uuid.UUID) (*models.StudentMarks, error) {
	var sm models.StudentMarks
	err := ac.dbConn.Model(sm).Where("student_id = ?", id).Find(&sm).Error
	if err != nil {
		return nil, err
	}
	return &sm, nil
}

func (ac *AdminstrationCloud) UpdateStudentMarks(sm *models.StudentMarks) error {

	err := ac.dbConn.Save(&sm).Error

	if err != nil {
		return err
	}
	return nil

}
