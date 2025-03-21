package daos

import (
	"CollegeAdministration/models"

	"github.com/google/uuid"
)

func (dao *Daos) CreateStudentMarks(sm *models.StudentMarks) error {

	err := dao.dbConn.Table("student_marks").Create(sm).Error

	if err != nil {
		return err
	}
	return nil
}

func (dao *Daos) GetMarksByMarksId(id uuid.UUID) (*models.StudentMarks, error) {

	var sm *models.StudentMarks
	err := dao.dbConn.Model(sm).Where("id = ?", id).Find(&sm).Error

	if err != nil {
		return nil, err
	}
	return sm, nil
}
func (dao *Daos) GetMarksByStudentId(id uuid.UUID) (*models.StudentMarks, error) {
	var sm models.StudentMarks
	err := dao.dbConn.Model(sm).Where("student_id = ?", id).Find(&sm).Error
	if err != nil {
		return nil, err
	}
	return &sm, nil
}

func (dao *Daos) UpdateStudentMarks(sm *models.StudentMarks) error {

	err := dao.dbConn.Table("student_marks").Where("id = ?", sm.Id).Updates(models.StudentMarks{Marks: sm.Marks, Grade: sm.Grade}).Error

	if err != nil {
		return err
	}
	return nil

}

func (dao *Daos) GetAllStudentsIDForACourse(course_id uuid.UUID) ([]string, error) {
	var student_names []string
	err := dao.dbConn.Table("student_marks").Select("student_id").Where("course_id = ?", course_id).Order("marks desc").Find(&student_names).Error
	if err != nil {
		return nil, err
	}

	return student_names, nil
}

func (dao *Daos) DeleteStudenetMarks(marks_id uuid.UUID) error {

	err := dao.dbConn.Table("student_marks").Where("id = ?", marks_id).Delete(models.StudentMarks{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *Daos) UpdateStudentMarksTableCourse(new_course string, student_id uuid.UUID) error {

	err := dao.dbConn.Table("student_marks").Where("student_id = ?", student_id).Update("course_name", new_course).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *Daos) GetMarksByStudentIdAndCourseId(student_id uuid.UUID, course_id uuid.UUID) (*models.StudentMarks, error) {
	var sm models.StudentMarks
	err := dao.dbConn.Model(sm).Where("student_id = ? and course_id = ?", student_id, course_id).Find(&sm).Error
	if err != nil {
		return nil, err
	}
	return &sm, nil
}
