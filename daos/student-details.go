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
		return fmt.Errorf("failed! %s", err.Error())
	}
	return nil

}
func (ac *AdminstrationCloud) RetieveCollegeAdminstration() ([]*models.StudentInfo, error) {

	var rca []*models.StudentInfo
	err := ac.dbConn.Order("roll_number").Order("name").Find(&rca).Error
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
		existingRC, err := ac.GetMarksByMarksId(eachRCA.MarksId)
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

func (ac *AdminstrationCloud) RetieveCollegeAdminstrationByOrder(order_by string) ([]*models.StudentInfo, error) {

	var rca []*models.StudentInfo
	if order_by == "roll_number" || order_by == "age" || order_by == "name" {
		err := ac.dbConn.Order(order_by).Find(&rca).Error
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
			existingRC, err := ac.GetMarksByMarksId(eachRCA.MarksId)
			if existingRC.Id == uuid.Nil {
				continue
			} else if err != nil {
				return nil, err
			} else {
				eachRCA.StudentMarks = *existingRC
			}
		}
	}
	return rca, nil

}

func (ac *AdminstrationCloud) UpdateClgStudent(rca *models.StudentInfo) error {

	err := ac.dbConn.Save(&rca).Error

	if err != nil {
		return fmt.Errorf("failed to UpdateClgStudent %s", err.Error())
	}
	return nil
}

func (ac *AdminstrationCloud) GetStudentDetailsByRollNumber(roll_number string) ([]*models.StudentInfo, error) {

	var cad []*models.StudentInfo
	val := ac.dbConn.Select("*").Table("student_infos").Where("roll_number = ?", roll_number).First(&cad)
	if val.Error != nil {
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

func (ac *AdminstrationCloud) GetStudentDetailsByRollNumberAndCourseId(roll_number string, courseId uuid.UUID) (*models.StudentInfo, error) {

	var cad *models.StudentInfo
	val := ac.dbConn.Select("*").Table("student_infos").Where("roll_number = ? AND course_id = ?", roll_number, courseId).Find(&cad)
	if val.Error != nil {
		log.Println("Not able to Fetch values student_infos table ", val.Error)
		return nil, val.Error
	}
	return cad, nil
}

func (ac *AdminstrationCloud) DeleteCourseForAStudent(st_name string, c_id uuid.UUID) error {

	err := ac.dbConn.Where("name = ? AND course_id = ?", st_name, c_id).Delete(models.StudentInfo{}).Error
	if err != nil {
		return err
	}
	return nil
}
func (ac *AdminstrationCloud) GetStudentdetail(sd *models.StudentInfo) (*models.StudentInfo, error) {

	var sd1 models.StudentInfo
	condition := make(map[string]interface{})
	if sd.Id != uuid.Nil {
		condition["id"] = sd.Id
	}
	if sd.Name != "" {
		condition["name"] = sd.Name
	}
	if sd.Age != 0 {
		condition["age"] = sd.Age
	}
	if sd.RollNumber != "" {
		condition["roll_number"] = sd.RollNumber
	}

	if sd.CourseId != uuid.Nil {
		condition["course_id"] = sd.CourseId
	}

	err := ac.dbConn.Model(models.StudentInfo{}).Where(condition).Find(&sd1).Error
	if err != nil {
		return nil, err
	}
	if sd1.Id == uuid.Nil {
		return nil, fmt.Errorf("no record found")
	}
	return &sd1, nil
}
func (ac *AdminstrationCloud) DeleteStudentWithSpecifics(st_req *models.StudentInfo) error {

	err := ac.dbConn.Where(st_req).Delete(models.StudentInfo{}).Error
	if err != nil {
		return err
	}
	return nil
}
