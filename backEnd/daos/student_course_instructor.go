package daos

import "CollegeAdministration/models"

func (dao *Daos) InsertIntoSCI(studentCourseInstructor *models.StudentCourseInstructor) error {

	err := dao.dbConn.Create(studentCourseInstructor).Error

	if err != nil {
		return err
	}
	return nil
}

func (dao *Daos) DeleteStudentCourseInstructor(studentId string, courseId string, instructorId string) error {
	err := dao.dbConn.Where("student_id = ? AND course_id = ? AND instructor_id = ?", studentId, courseId, instructorId).Update("is_deleted", true).Error
	if err != nil {
		return err
	}
	return nil
}
func (dao *Daos) HardDeleteStudentCourseInstructor(studentId string, courseId string, instructorId string) error {
	err := dao.dbConn.Where("student_id = ? AND course_id = ? AND instructor_id = ?", studentId, courseId, instructorId).Delete(models.StudentCourseInstructor{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *Daos) HardDeleteStudentCourseInstructorByStudentId(studentId string) error {
	err := dao.dbConn.Where("student_id = ?", studentId).Delete(models.StudentCourseInstructor{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *Daos) GetSCIByStudentIdAndCourseIdAndInstructorId(studentId string, courseId string, instructorId string) (*models.StudentCourseInstructor, error) {
	var studentCourseInstructor *models.StudentCourseInstructor
	err := dao.dbConn.Where("student_id = ? AND course_id = ? AND instructor_id = ?", studentId, courseId, instructorId).First(&studentCourseInstructor).Error
	if err != nil {
		return nil, err
	}
	return studentCourseInstructor, nil
}
func (dao *Daos) GetSCIByStudentIdAndCourseId(studentId string, courseId string) ([]*models.StudentCourseInstructor, error) {
	var studentCourseInstructor []*models.StudentCourseInstructor
	err := dao.dbConn.Where("student_id = ? AND course_id = ? AND is_deleted = false", studentId, courseId).Find(&studentCourseInstructor).Error
	if err != nil {
		return nil, err
	}
	return studentCourseInstructor, nil
}
func (dao *Daos) GetSCIByStudentId(studentId string) ([]*models.StudentCourseInstructor, error) {
	var studentCourseInstructor []*models.StudentCourseInstructor
	err := dao.dbConn.Where("student_id = ? AND is_deleted = false", studentId).Find(&studentCourseInstructor).Error
	if err != nil {
		return nil, err
	}
	return studentCourseInstructor, nil
}
func (dao *Daos) GetSCIByCourseId(courseId string) ([]*models.StudentCourseInstructor, error) {
	var studentCourseInstructor []*models.StudentCourseInstructor
	err := dao.dbConn.Where("course_id = ? AND is_deleted = false", courseId).Find(&studentCourseInstructor).Error
	if err != nil {
		return nil, err
	}
	return studentCourseInstructor, nil
}
func (dao *Daos) GetSCIByInstructorId(instructorId string) ([]*models.StudentCourseInstructor, error) {
	var studentCourseInstructor []*models.StudentCourseInstructor
	err := dao.dbConn.Where("instructor_id = ? AND is_deleted = false", instructorId).Find(&studentCourseInstructor).Error
	if err != nil {
		return nil, err
	}
	return studentCourseInstructor, nil
}
func (dao *Daos) GetSCIByInstructorIdAndCourseId(instructorId string, courseId string) ([]*models.StudentCourseInstructor, error) {
	var studentCourseInstructor []*models.StudentCourseInstructor
	err := dao.dbConn.Where("instructor_id = ? AND course_id = ? AND is_deleted = false", instructorId, courseId).Find(&studentCourseInstructor).Error
	if err != nil {
		return nil, err
	}
	return studentCourseInstructor, nil
}
func (dao *Daos) GetSCIByInstructorIdAndStudentId(instructorId string, studentId string) ([]*models.StudentCourseInstructor, error) {
	var studentCourseInstructor []*models.StudentCourseInstructor
	err := dao.dbConn.Where("instructor_id = ? AND student_id = ? AND is_deleted = false", instructorId, studentId).Find(&studentCourseInstructor).Error
	if err != nil {
		return nil, err
	}
	return studentCourseInstructor, nil
}
func (dao *Daos) GetSCIByCourseIdAndStudentId(courseId string, studentId string) ([]*models.StudentCourseInstructor, error) {
	var studentCourseInstructor []*models.StudentCourseInstructor
	err := dao.dbConn.Where("course_id = ? AND student_id = ? AND is_deleted = false", courseId, studentId).Find(&studentCourseInstructor).Error
	if err != nil {
		return nil, err
	}
	return studentCourseInstructor, nil
}
func (dao *Daos) GetSCIByCourseIdAndInstructorId(courseId string, instructorId string) ([]*models.StudentCourseInstructor, error) {
	var studentCourseInstructor []*models.StudentCourseInstructor
	err := dao.dbConn.Where("course_id = ? AND instructor_id = ? AND is_deleted = false", courseId, instructorId).Find(&studentCourseInstructor).Error
	if err != nil {
		return nil, err
	}
	return studentCourseInstructor, nil
}
func (dao *Daos) GetSCIByStudentIdAndInstructorId(studentId string, instructorId string) ([]*models.StudentCourseInstructor, error) {
	var studentCourseInstructor []*models.StudentCourseInstructor
	err := dao.dbConn.Where("student_id = ? AND instructor_id = ? AND is_deleted = false", studentId, instructorId).Find(&studentCourseInstructor).Error
	if err != nil {
		return nil, err
	}
	return studentCourseInstructor, nil
}
func (dao *Daos) GetSCIByStudentIdAndCourseIdAndInstructorIdAndIsDeleted(studentId string, courseId string, instructorId string, isDeleted bool) (*models.StudentCourseInstructor, error) {
	var studentCourseInstructor *models.StudentCourseInstructor
	err := dao.dbConn.Where("student_id = ? AND course_id = ? AND instructor_id = ? AND is_deleted = ?", studentId, courseId, instructorId, isDeleted).First(&studentCourseInstructor).Error
	if err != nil {
		return nil, err
	}
	return studentCourseInstructor, nil
}
func (dao *Daos) GetSCIByStudentIdAndCourseIdAndInstructorIdAndIsDeletedAndIsCompleted(studentId string, courseId string, instructorId string, isDeleted bool, isCompleted bool) (*models.StudentCourseInstructor, error) {
	var studentCourseInstructor *models.StudentCourseInstructor
	err := dao.dbConn.Where("student_id = ? AND course_id = ? AND instructor_id = ? AND is_deleted = ? AND is_completed = ?", studentId, courseId, instructorId, isDeleted, isCompleted).First(&studentCourseInstructor).Error
	if err != nil {
		return nil, err
	}
	return studentCourseInstructor, nil
}
func (dao *Daos) GetSCIByStudentIdAndCourseIdAndInstructorIdAndIsDeletedAndIsCompletedAndIsGraded(studentId string, courseId string, instructorId string, isDeleted bool, isCompleted bool, isGraded bool) (*models.StudentCourseInstructor, error) {
	var studentCourseInstructor *models.StudentCourseInstructor
	err := dao.dbConn.Where("student_id = ? AND course_id = ? AND instructor_id = ? AND is_deleted = ? AND is_completed = ? AND is_graded = ?", studentId, courseId, instructorId, isDeleted, isCompleted, isGraded).First(&studentCourseInstructor).Error
	if err != nil {
		return nil, err
	}
	return studentCourseInstructor, nil
}
