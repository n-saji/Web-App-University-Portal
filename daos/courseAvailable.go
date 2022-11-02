package daos

import (
	"CollegeAdministration/models"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (ac *AdminstrationCloud) InsertValuesToCoursesAvailable(ca *models.CoursesAvailable) error {

	err := ac.dbConn.Table("courses_availables").Create(ca).Error
	if err != nil {
		log.Println("Not able to insert to courses_availables table ", err)
		return fmt.Errorf("Failed! ", err)
	}
	log.Println("Stored to database")
	return nil

}
func (ac *AdminstrationCloud) GetCourseByName(name string) (models.CoursesAvailable, error) {

	var ca models.CoursesAvailable
	val := ac.dbConn.Select("*").Table("courses_availables").Where("course_name = ?", name).First(&ca)
	if val.Error != nil {
		log.Println("Not able to Fetch value from  table ", val.Error)
		return ca, val.Error
	}
	return ca, nil
}

func (ac *AdminstrationCloud) GetCourseById(id uuid.UUID) (models.CoursesAvailable, error) {

	var ca models.CoursesAvailable
	val := ac.dbConn.Select("*").Table("courses_availables").Where("id = ?", id).First(&ca)
	if val.Error != nil {
		log.Println("Not able to insert to courses_availables table ", val.Error)
		return ca, val.Error
	}
	return ca, nil
}

func (ac *AdminstrationCloud) RetieveCoursesAvailable() ([]*models.CoursesAvailable, error) {

	var rca []*models.CoursesAvailable
	err := ac.dbConn.Find(&rca).Error

	return rca, err

}
func (ac *AdminstrationCloud) UpdateCourseByName(name string, rc *models.CoursesAvailable) error {
	rcOld, _ := ac.GetCourseByName(name)

	rc.Id = rcOld.Id
	err := ac.dbConn.Save(&rc).Error
	if err != nil {
		return fmt.Errorf("Failed UpdateCourseByName ", err)
	}
	return nil
}

func (ac *AdminstrationCloud) CheckCourse(coursename string) bool {

	var len int64
	ac.dbConn.Model(models.CoursesAvailable{}).Where("course_name = ?", coursename).Count(&len)
	if len > 0 {
		return true
	} else {
		return false
	}
}
