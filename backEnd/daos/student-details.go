package daos

import (
	"CollegeAdministration/models"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (dao *Daos) InsertIntoStudentInfos(ca *models.StudentInfo) error {

	err := dao.dbConn.Table("student_infos").Create(ca).Error
	if err != nil {
		log.Println("Not able to insert to student_infos table ", err)
		return fmt.Errorf("failed! %s", err.Error())
	}
	return nil

}

func (dao *Daos) RetrieveCollegeAdministrationByOrder(order_by string) ([]*models.StudentInfo, error) {

	var rca []*models.StudentInfo
	if order_by == "roll_number" || order_by == "age" || order_by == "name" {
		err := dao.dbConn.Order(order_by).Find(&rca).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := dao.dbConn.Find(&rca).Error
		if err != nil {
			return nil, err
		}
	}

	// if order_by == "course_name" || order_by == "marks" || order_by == "grade" {
	// 	if order_by == "course_name" {
	// 		for i := 0; i < len(rca)-1; i++ {
	// 			for j := i; j < len(rca); j++ {
	// 				if rca[i].ClassesEnrolled.CourseName > rca[j].ClassesEnrolled.CourseName {
	// 					temp := rca[j]
	// 					rca[j] = rca[i]
	// 					rca[i] = temp
	// 				}
	// 			}
	// 		}
	// 	}
	// 	if order_by == "grade" {
	// 		for i := 0; i < len(rca)-1; i++ {
	// 			for j := i; j < len(rca); j++ {
	// 				if rca[i].StudentMarks.Grade > rca[j].StudentMarks.Grade {
	// 					temp := rca[j]
	// 					rca[j] = rca[i]
	// 					rca[i] = temp
	// 				}
	// 			}
	// 		}
	// 	}
	// 	if order_by == "marks" {
	// 		for i := 0; i < len(rca)-1; i++ {
	// 			for j := i; j < len(rca); j++ {
	// 				if rca[i].StudentMarks.Marks > rca[j].StudentMarks.Marks {
	// 					temp := rca[j]
	// 					rca[j] = rca[i]
	// 					rca[i] = temp
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	return rca, nil

}

func (dao *Daos) UpdateClgStudent(rca *models.StudentInfo) error {

	err := dao.dbConn.Save(&rca).Error

	if err != nil {
		return fmt.Errorf("failed to UpdateClgStudent %s", err.Error())
	}
	return nil
}

func (dao *Daos) GetStudentDetailsByRollNumber(roll_number string) (*models.StudentInfo, error) {

	var cad *models.StudentInfo
	val := dao.dbConn.Select("*").Table("student_infos").Where("roll_number = ?", roll_number).First(&cad)
	if val.Error != nil {
		return nil, val.Error
	}
	return cad, nil
}
func (dao *Daos) CheckForRollNo(roll_number string) (bool, error) {

	var len int64
	err := dao.dbConn.Model(models.StudentInfo{}).Where("roll_number = ?", roll_number).Count(&len).Error
	if err != nil {
		return false, err
	}

	if len > 0 {
		return true, nil
	} else {
		return false, nil
	}

}

func (dao *Daos) GetStudentdetailsUsingCourseId(courseId uuid.UUID) ([]*models.StudentInfo, error) {

	var rca []*models.StudentInfo

	err := dao.dbConn.Select("*").Table("student_infos").Where("course_id = ?", courseId).Find(&rca).Error
	if err != nil {
		return nil, nil
	}

	return rca, nil
}

func (dao *Daos) DeleteStudentDaos(studentId uuid.UUID) error {

	err := dao.dbConn.Where("id = ?", studentId).Delete(&models.StudentInfo{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (dao *Daos) GetStudentDetailsByName(student_name string) (*[]models.StudentInfo, error) {

	var si *[]models.StudentInfo

	err := dao.dbConn.Select("*").Table("student_infos").Where("name = ?", student_name).Find(&si).Error

	if err != nil {
		return nil, err
	}
	if len(*si) == 0 {
		return nil, fmt.Errorf("no student exists")
	}
	return si, nil
}

func (dao *Daos) GetStudentDetailsByRollNumberAndCourseId(roll_number string, courseId uuid.UUID) (*models.StudentInfo, error) {

	var cad *models.StudentInfo
	val := dao.dbConn.Select("*").Table("student_infos").Where("roll_number = ? AND course_id = ?", roll_number, courseId).Find(&cad)
	if val.Error != nil {
		log.Println("Not able to Fetch values student_infos table ", val.Error)
		return nil, val.Error
	}
	return cad, nil
}

func (dao *Daos) DeleteCourseForAStudent(st_name string, c_id uuid.UUID) error {

	err := dao.dbConn.Where("name = ? AND course_id = ?", st_name, c_id).Delete(models.StudentInfo{}).Error
	if err != nil {
		return err
	}
	return nil
}
func (dao *Daos) GetStudentdetail(sd *models.StudentInfo) (*models.StudentInfo, error) {

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

	// if sd.CourseId != uuid.Nil {
	// 	condition["course_id"] = sd.CourseId
	// }

	err := dao.dbConn.Model(models.StudentInfo{}).Where(condition).Find(&sd1).Error
	if err != nil {
		return nil, err
	}
	if sd1.Id == uuid.Nil {
		return nil, fmt.Errorf("no record found")
	}
	return &sd1, nil
}
func (dao *Daos) DeleteStudentWithSpecifics(st_req *models.StudentInfo) error {

	err := dao.dbConn.Where(st_req).Delete(models.StudentInfo{}).Error
	if err != nil {
		return err
	}
	return nil
}
