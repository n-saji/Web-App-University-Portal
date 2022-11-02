package daos

import (
	"CollegeAdministration/models"
	"fmt"
	"log"
<<<<<<< HEAD
=======

	"github.com/google/uuid"
>>>>>>> feature_branch
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
<<<<<<< HEAD

	for _, eachRCA := range rca {
		eachRCA.ClassesEnrolled, err = ac.GetCourseById(eachRCA.CourseId)
		if err != nil {
			return rca, err
		}
	}
	return rca, err
=======
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
>>>>>>> feature_branch

}

func (ac *AdminstrationCloud) UpdateClgStudent(rca *models.CollegeAdminstration) error {

<<<<<<< HEAD
	// rcaExisting, _ := ac.GetStudentDetailsByRollNumber(rca.RollNumber)
	// rcaExisting.Name = rca.Name
	// rcaExisting.Age = rca.Age
	// rcaExisting.ClassesEnrolled.Id = rca.ClassesEnrolled.Id
	// rcaExisting.ClassesEnrolled.CourseName = rca.ClassesEnrolled.CourseName
	// log.Println(rcaExisting)
	// err := ac.dbConn.Save(&rcaExisting).Error
	//err := ac.dbConn.Model(&models.CollegeAdminstration{}).Where("Id = ?", rca.Id).Updates(map[string]interface{}{"Name": rca.Name, "Age": rca.Age}) //Save(&rca).Error
=======
>>>>>>> feature_branch
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
<<<<<<< HEAD
	//log.Println(len)
=======
>>>>>>> feature_branch

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
