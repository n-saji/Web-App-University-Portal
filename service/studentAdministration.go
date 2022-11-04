package service

import (
	"CollegeAdministration/models"
	"fmt"

	"github.com/google/uuid"
)

func (ac *Service) InsertValuesToCAd(cv *models.StudentInfo) error {

	cv_id, err := ac.daos.GetCourseByName(cv.ClassesEnrolled.CourseName)
	if err != nil {
		return fmt.Errorf("Course Not Found")
	}
	cv.ClassesEnrolled.Id = cv_id.Id
	/* ok, err := ac.daos.CheckForRollNo(cv.RollNumber)
	if err != nil {
		return err
	} else {
		if ok {
			return fmt.Errorf("Roll Number already exist!")
		}
	} */

	cv.Id = uuid.New()
	err1 := ac.daos.InsertValuesToCollegeAdminstration(cv)
	if err1 != nil {
		return err1
	} else {
		return nil
	}

}

func (ac *Service) RetrieveCAd() ([]*models.StudentInfo, error) {

	rca, err := ac.daos.RetieveCollegeAdminstration()
	if err != nil {
		return rca, err
	}
	return rca, nil
}

func (ac *Service) UpdateCAd(rca *models.StudentInfo) error {
	rc, err1 := ac.daos.GetCourseByName(rca.ClassesEnrolled.CourseName)
	if err1 != nil {
		return fmt.Errorf("Course not Found")
		//rc.CourseName = rca.ClassesEnrolled.CourseName
		//ac.InsertValuesToCA(&rc)
	}

	if rca.Id == uuid.Nil {
		rcaOld, err := ac.daos.GetStudentDetailsByRollNumber(rca.RollNumber)
		if err != nil {
			return fmt.Errorf("ROLL number not found", err)
		}
		rca.Id = rcaOld.Id
	}
	if rca.ClassesEnrolled.Id == uuid.Nil {
		rc, _ = ac.daos.GetCourseByName(rca.ClassesEnrolled.CourseName)
		rca.ClassesEnrolled.Id = rc.Id
		rca.CourseId = rc.Id
	}
	err := ac.daos.UpdateClgStudent(rca)
	if err != nil {
		return err
	}
	return nil

}

func (ac *Service) DeleteStudent(rollNumber string) error {

	student, err := ac.daos.GetStudentDetailsByRollNumber(rollNumber)
	if err != nil {
		return err
	}
	err1 := ac.daos.DeleteStudentDaos(student.Id)
	if err1 != nil {
		return err1
	}
	return nil
}

func (ac *Service) UpdateStudentNameAge(existing_name, student_name string, age int64) error {
	si, err := ac.daos.GetStudentDetailsByName(existing_name)
	if err != nil {
		return err
	}
	for _, each_si := range *si {
		each_si.Name = student_name
		if age != 0 {
			each_si.Age = age
		}
		err := ac.daos.UpdateClgStudent(&each_si)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ac *Service) FetchStudentCourse(student_name string) (map[string]string, error) {
	course_list := make(map[string]string)

	course_list["student_name"] = student_name
	si, err := ac.daos.GetStudentDetailsByName(student_name)
	if err != nil {
		return nil, err
	}
	for index, each_si := range *si {
		each_si.ClassesEnrolled, err = ac.daos.GetCourseById(each_si.CourseId)
		if err != nil {
			return nil, err
		}
		course_list[fmt.Sprintf("course_%d", index+1)] = each_si.ClassesEnrolled.CourseName
	}
	return course_list, nil
}
