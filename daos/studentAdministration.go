package daos

import (
	"CollegeAdministration/models"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (ac *AdminstrationCloud) InsertValuesToCollegeAdminstration(ca *models.StudentInfo) error {

	err := ac.dbConn.Table("student_infos").Create(ca).Error
	if err != nil {
		log.Println("Not able to insert to student_infos table ", err)
		return fmt.Errorf("Failed! ", err)
	}
	log.Println("Stored to database")
	return nil

}
func (ac *AdminstrationCloud) RetieveCollegeAdminstration() ([]*models.StudentInfo, error) {

	var rca []*models.StudentInfo
	err := ac.dbConn.Find(&rca).Error
	if err != nil {
		return nil, err
	}

	for _, eachRCA := range rca {
		existingRC, err := ac.GetCourseById(eachRCA.CourseId)
		if existingRC.Id == uuid.Nil {
			continue
		} else if err != nil {
			return nil, err
		} else {
			eachRCA.ClassesEnrolled = existingRC
		}
	}

	for _, eachRCA := range rca {
		existingRC, err := ac.GetMarksById(eachRCA.MarksId)
		if existingRC.Id == uuid.Nil {
			continue
		} else if err != nil {
			return nil, err
		} else {
			eachRCA.StudentMarks = *existingRC
		}
	}
	return rca, nil

}

func (ac *AdminstrationCloud) UpdateClgStudent(rca *models.StudentInfo) error {

	err := ac.dbConn.Save(&rca).Error

	if err != nil {
		return fmt.Errorf("Failed to UpdateClgStudent", err)
	}
	return nil
}

func (ac *AdminstrationCloud) GetStudentDetailsByRollNumber(roll_number string) (*models.StudentInfo, error) {

	var cad *models.StudentInfo
	val := ac.dbConn.Select("*").Table("student_infos").Where("roll_number = ?", roll_number).First(&cad)
	if val.Error != nil {
		log.Println("Not able to insert to student_infos table ", val.Error)
		return nil, val.Error
	}
	return cad, nil
}
func (ac *AdminstrationCloud) CheckForRollNo(roll_number string) (bool, error) {

	var len int64
	err := ac.dbConn.Model(models.StudentInfo{}).Where("roll_number = ?", roll_number).Count(&len).Error
	if err != nil {
		return false, err
	}

	if len > 0 {
		return true, nil
	} else {
		return false, nil
	}

}

func (ac *AdminstrationCloud) GetStudentdetailsUsingCourseId(courseId uuid.UUID) ([]*models.StudentInfo, error) {

	var rca []*models.StudentInfo

	err := ac.dbConn.Select("*").Table("student_infos").Where("course_id = ?", courseId).Find(&rca).Error
	if err != nil {
		return nil, nil
	}

	return rca, nil
}

func (ac *AdminstrationCloud) DeleteStudentDaos(studentId uuid.UUID) error {

	err := ac.dbConn.Where("id = ?", studentId).Delete(&models.StudentInfo{}).Error

	if err != nil {
		return nil
	}

	return nil
}

func (ac *AdminstrationCloud) GetStudentDetailsByName(student_name string) (*[]models.StudentInfo, error) {

	var si *[]models.StudentInfo

	err := ac.dbConn.Select("*").Table("student_infos").Where("name = ?", student_name).Find(&si).Error

	if err != nil {
		return nil, err
	}
	if len(*si) == 0 {
		return nil, fmt.Errorf("no student exists")
	}
	return si, nil
}
