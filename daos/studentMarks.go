package daos

import (
	"CollegeAdministration/models"

	"github.com/google/uuid"
)

func (ac *AdministrationCloud) CreateStudentMarks(sm *models.StudentMarks) error {

	err := ac.dbConn.Table("student_marks").Create(sm).Error

	if err != nil {
		return err
	}
	return nil
}

func (ac *AdministrationCloud) GetMarksByMarksId(id uuid.UUID) (*models.StudentMarks, error) {

	var sm *models.StudentMarks
	err := ac.dbConn.Model(sm).Where("id = ?", id).Find(&sm).Error

	if err != nil {
		return nil, err
	}
	return sm, nil
}
func (ac *AdministrationCloud) GetMarksByStudentId(id uuid.UUID) (*models.StudentMarks, error) {
	var sm models.StudentMarks
	err := ac.dbConn.Model(sm).Where("student_id = ?", id).Find(&sm).Error
	if err != nil {
		return nil, err
	}
	return &sm, nil
}

func (ac *AdministrationCloud) UpdateStudentMarks(sm *models.StudentMarks) error {

	err := ac.dbConn.Table("student_marks").Where("id = ?", sm.Id).Updates(models.StudentMarks{Marks: sm.Marks, Grade: sm.Grade}).Error 

	if err != nil {
		return err
	}
	return nil

}

func (ac *AdministrationCloud) GetAllStudentsIDForACourse(course_id uuid.UUID) ([]string, error) {
	var student_names []string
	err := ac.dbConn.Table("student_marks").Select("student_id").Where("course_id = ?", course_id).Order("marks desc").Find(&student_names).Error
	if err != nil {
		return nil, err
	}

	return student_names, nil
}

func (ac *AdministrationCloud) DeleteStudenetMarks(marks_id uuid.UUID) error {

	err := ac.dbConn.Table("student_marks").Where("id = ?", marks_id).Delete(models.StudentMarks{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (ac *AdministrationCloud) UpdateStudentMarksTableCourse(new_course string,student_id uuid.UUID)error{

	err := ac.dbConn.Table("student_marks").Where("student_id = ?",student_id).Update("course_name",new_course).Error
	if err != nil {
		return err
	}
	return nil
}