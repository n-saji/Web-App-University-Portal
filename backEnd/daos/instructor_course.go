package daos

import "CollegeAdministration/models"

func (dao *Daos) InsertIntoIC(instructorCourse *models.InstructorCourse) error {

	err := dao.dbConn.Create(instructorCourse).Error

	if err != nil {
		return err
	}
	return nil
}

func (dao *Daos) GetCoursesByInstructorId(instructorId string) ([]*models.InstructorCourse, error) {
	var instructorCourse []*models.InstructorCourse
	err := dao.dbConn.Where("instructor_id = ? and is_deleted = false", instructorId).Find(&instructorCourse).Error
	if err != nil {
		return nil, err
	}
	return instructorCourse, nil
}

func (dao *Daos) GetAllCoursesByInstructorId(instructorId string) ([]*models.InstructorCourse, error) {
	var instructorCourse []*models.InstructorCourse
	err := dao.dbConn.Where("instructor_id = ?", instructorId).Find(&instructorCourse).Error
	if err != nil {
		return nil, err
	}
	return instructorCourse, nil
}

func (dao *Daos) GetInstructorsByCourseId(courseId string) ([]*models.InstructorCourse, error) {
	var instructorCourse []*models.InstructorCourse
	err := dao.dbConn.Where("course_id = ? and is_deleted = false", courseId).Find(&instructorCourse).Error
	if err != nil {
		return nil, err
	}
	return instructorCourse, nil
}

func (dao *Daos) DeleteInstructorCourse(instructorId string, courseId string) error {
	err := dao.dbConn.Where("instructor_id = ? AND course_id = ?", instructorId, courseId).Update("is_deleted", true).Error
	if err != nil {
		return err
	}
	return nil
}
func (dao *Daos) HardDeleteInstructorCourse(instructorId string, courseId string) error {
	err := dao.dbConn.Where("instructor_id = ? AND course_id = ?", instructorId, courseId).Delete(models.InstructorCourse{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *Daos) GetICByInstructorIdAndCourseId(instructorId string, courseId string) (*models.InstructorCourse, error) {
	var instructorCourse *models.InstructorCourse
	err := dao.dbConn.Where("instructor_id = ? AND course_id = ?", instructorId, courseId).First(&instructorCourse).Error
	if err != nil {
		return nil, err
	}
	return instructorCourse, nil
}

func (dao *Daos) UpdateInstructorCourse(instructorCourse *models.InstructorCourse) error {
	err := dao.dbConn.Model(models.InstructorCourse{}).Where("instructor_id = ? AND course_id = ?", instructorCourse.InstructorId, instructorCourse.CourseId).Updates(instructorCourse).Error
	if err != nil {
		return err
	}
	return nil
}
