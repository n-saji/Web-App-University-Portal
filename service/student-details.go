package service

import (
	"CollegeAdministration/models"
	"fmt"

	"github.com/google/uuid"
)

func (ac *Service) InsertValuesToCAd(cv *models.StudentInfo) error {

	cv_id, err := ac.daos.GetCourseByName(cv.ClassesEnrolled.CourseName)
	if err != nil {
		return fmt.Errorf("course Not Found")
	}

	cv.ClassesEnrolled.Id = cv_id.Id
	cv.CourseId = cv_id.Id
	sd_existing, _ := ac.daos.GetStudentDetailsByRollNumber(cv.RollNumber)
	if sd_existing != nil && sd_existing.Name != cv.Name {
		return fmt.Errorf("roll number exits for another student")
	}
	sd, _ := ac.daos.GetStudentdetail(cv)
	if sd != nil && sd.CourseId == cv.CourseId {
		return fmt.Errorf("student with course exist")
	}

	cv.Id = uuid.New()
	sm, err2 := ac.InsertStudentIdInToMarksTable(cv)
	if err2 != nil {
		return err2
	}
	cv.MarksId = sm.Id

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

func (ac *Service) UpdateCAd(rca *models.StudentInfo, oldCourse string) error {

	rcOld, err4 := ac.daos.GetCourseByName(oldCourse)
	rcNew, err1 := ac.daos.GetCourseByName(rca.ClassesEnrolled.CourseName)
	if err1 != nil {
		return fmt.Errorf("%s Course not Found", rcNew.CourseName)
	}
	if err4 != nil {
		return fmt.Errorf("%s Course not Found", oldCourse)
	}
	rcaExist, _ := ac.daos.GetStudentdetail(
		&models.StudentInfo{
			RollNumber: rca.RollNumber,
			Id:         rcOld.Id,
			Name:       rca.Name})
	if rcaExist == nil {
		return fmt.Errorf("student details mismatched")
	}

	if rcaExist.Id == uuid.Nil {
		return fmt.Errorf("the student does not have %s registered", rcOld.CourseName)
	}
	rcaNew, _ := ac.daos.GetStudentDetailsByRollNumberAndCourseId(rca.RollNumber, rcNew.Id)
	if rcaNew.Id != uuid.Nil && rcOld.Id != rcNew.Id {
		return fmt.Errorf("the student already has %s registered can't duplicate course ,please make course in url and body same", rcNew.CourseName)
	}

	if rcOld.Id != rcNew.Id && rcaNew.Id == uuid.Nil {
		rca.Id = rcaExist.Id

	} else {

		rcaOld, err := ac.daos.GetStudentDetailsByRollNumberAndCourseId(rca.RollNumber, rcOld.Id)
		if err != nil {
			return fmt.Errorf("student roll number not found %s", err.Error())
		}
		rca.Id = rcaOld.Id
	}

	sm, err2 := ac.daos.GetMarksByStudentId(rca.Id)
	if err2 != nil {
		return err2
	}
	if rca.ClassesEnrolled.Id == uuid.Nil {
		rcNew, _ = ac.daos.GetCourseByName(rca.ClassesEnrolled.CourseName)
		rca.ClassesEnrolled.Id = rcNew.Id
		rca.CourseId = rcNew.Id
	}
	rca.MarksId = sm.Id
	sm.Marks = rca.StudentMarks.Marks
	sm.Grade = ac.GenerateGradeForMarks(sm.Marks)
	sm.CourseName = rca.ClassesEnrolled.CourseName
	sm.CourseId = rca.CourseId
	rca.StudentMarks = *sm

	err3 := ac.daos.UpdateStudentMarks(sm)
	if err3 != nil {
		return err3
	}
	err5 := ac.daos.UpdateClgStudent(rca)
	if err5 != nil {
		return err5
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

		studentmarks, err := ac.daos.GetMarksByStudentId(each_si.Id)
		each_si.StudentMarks = *studentmarks
		if err != nil {
			return nil, err
		}
		course_list[fmt.Sprintf("course_%d", index+1)] = fmt.Sprintf("%s -->Marks: %d Grade: %s", each_si.ClassesEnrolled.CourseName, each_si.StudentMarks.Marks, each_si.StudentMarks.Grade)
	}
	return course_list, nil
}

func (ac *Service) DeleteStudentCourseService(sn, cn string) (err error) {

	course_details, err1 := ac.daos.GetCourseByName(cn)
	if err1 != nil {
		return fmt.Errorf("course not found")
	}
	var student_detail models.StudentInfo
	student_detail.Name = sn
	student_detail.CourseId = course_details.Id
	_, err2 := ac.daos.GetStudentdetail(&student_detail)
	if err2 != nil {
		return err2
	}
	err = ac.daos.DeleteCourseForAStudent(sn, course_details.Id)
	if err != nil {
		return fmt.Errorf("failed to delete")
	}
	return nil

}

func (s *Service) GetAllStudentSelectiveData() ([]*models.StudentSelectiveData, error) {

	ssd := []*models.StudentSelectiveData{}

	student_data, err := s.daos.RetieveCollegeAdminstration()
	if err != nil {
		return nil, err
	}
	for _, each_student_data := range student_data {
		ssd = append(ssd,
			&models.StudentSelectiveData{
				Name:       each_student_data.Name,
				RollNumber: each_student_data.RollNumber,
				Course:     each_student_data.ClassesEnrolled.CourseName,
			})

	}

	return ssd, nil

}
