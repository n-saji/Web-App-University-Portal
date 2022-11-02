package daos

import (
	"CollegeAdministration/models"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (ac *AdminstrationCloud) InsertValuesToCollegeAdminstration(ca *models.CollegeAdminstration) error {

	err := ac.dbConn.Table("college_adminstrations").Create(ca).Error
	if err != nil {
		log.Println("Not able to insert to CollegeAdminstration table ", err)
		return fmt.Errorf("Failed! ", err)
	}
	log.Println("Stored to database")
	return nil

}
func (ac *AdminstrationCloud) RetieveCollegeAdminstration() ([]*models.CollegeAdminstration, error) {

	var rca []*models.CollegeAdminstration
	err := ac.dbConn.Find(&rca).Error
	if err != nil {
		return rca, err
	}

	for _, eachRCA := range rca {
		existingRC, err := ac.GetCourseById(eachRCA.CourseId)
		if existingRC.Id == uuid.Nil {
			continue
		} else if err != nil {
			return rca, err
		} else {
			eachRCA.ClassesEnrolled = existingRC
		}
	}
	return rca, nil

}

func (ac *AdminstrationCloud) UpdateClgStudent(rca *models.CollegeAdminstration) error {

	err := ac.dbConn.Save(&rca).Error

	if err != nil {
		return fmt.Errorf("Failed to UpdateClgStudent", err)
	}
	return nil
}

func (ac *AdminstrationCloud) GetStudentDetailsByRollNumber(roll_number string) (models.CollegeAdminstration, error) {

	var cad models.CollegeAdminstration
	val := ac.dbConn.Select("*").Table("college_adminstrations").Where("roll_number = ?", roll_number).First(&cad)
	if val.Error != nil {
		log.Println("Not able to insert to CollegeAdminstration table ", val.Error)
		return cad, val.Error
	}
	return cad, nil
}
func (ac *AdminstrationCloud) CheckForRollNo(roll_number string) (bool, error) {

	var len int64
	err := ac.dbConn.Model(models.CollegeAdminstration{}).Where("roll_number = ?", roll_number).Count(&len).Error
	if err != nil {
		return false, err
	}

	if len > 0 {
		return true, nil
	} else {
		return false, nil
	}

}

func (ac *AdminstrationCloud) GetStudentdetailsUsingCourseId(courseId uuid.UUID) ([]*models.CollegeAdminstration, error) {

	var rca []*models.CollegeAdminstration

	err := ac.dbConn.Select("*").Table("college_adminstrations").Where("course_id = ?", courseId).Find(&rca).Error
	if err != nil {
		return rca, nil
	}

	log.Println(rca)
	return rca, nil
}
