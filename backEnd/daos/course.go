package daos

import (
	"CollegeAdministration/models"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (dao *Daos) InsertValuesToCoursesAvailable(ca *models.CourseInfo) error {

	err := dao.dbConn.Table("course_infos").Create(ca).Error
	if err != nil {
		log.Println("not able to insert to course_infos table ", err)
		return fmt.Errorf("failed! %s", err.Error())
	}
	return nil

}
func (dao *Daos) GetCourseByName(name string) (models.CourseInfo, error) {

	var ca models.CourseInfo
	val := dao.dbConn.Select("*").Table("course_infos").Where("course_name = ?", name).First(&ca)
	if val.Error != nil {
		log.Println("Not able to Fetch value from  table ", val.Error)
		return ca, val.Error
	}
	return ca, nil
}

func (dao *Daos) GetCourseById(id uuid.UUID) (models.CourseInfo, error) {

	var ca models.CourseInfo
	// if id == uuid.Nil {
	// 	return ca, fmt.Errorf("UUID is NULL for course ID! Add new Course to ")
	// }
	val := dao.dbConn.Select("*").Table("course_infos").Where("id = ?", id).First(&ca)
	if val.Error != nil {
		log.Println("Not able to select from course_infos table ", val.Error)
		return ca, val.Error
	}
	return ca, nil
}

func (dao *Daos) RetieveCoursesAvailable() ([]*models.CourseInfo, error) {

	var rca []*models.CourseInfo
	err := dao.dbConn.Order("course_name").Find(&rca).Error

	return rca, err

}
func (dao *Daos) UpdateCourseByName(name string, rc *models.CourseInfo) error {
	rc_existing, _ := dao.GetCourseByName(rc.CourseName)
	if rc_existing.Id != uuid.Nil {
		return fmt.Errorf("course already exits")
	}
	rcOld, err1 := dao.GetCourseByName(name)
	if err1 != nil {
		return fmt.Errorf("course not available %s", err1.Error())
	}
	rc.Id = rcOld.Id
	err := dao.dbConn.Save(&rc).Error
	if err != nil {
		return fmt.Errorf("failed UpdateCourseByName %s", err.Error())
	}
	return nil
}

func (dao *Daos) CheckCourse(coursename string) bool {

	var len int64
	dao.dbConn.Model(models.CourseInfo{}).Where("course_name = ?", coursename).Count(&len)
	if len > 0 {
		return true
	} else {
		return false
	}
}

func (dao *Daos) DeleteCourse(id uuid.UUID) (bool, error) {

	err := dao.dbConn.Where("id = ?", id).Delete(&models.CourseInfo{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
