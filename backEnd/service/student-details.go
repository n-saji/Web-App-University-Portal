package service

import (
	"CollegeAdministration/config"
	"CollegeAdministration/models"
	"CollegeAdministration/utils"
	"fmt"
	"log"
	"strconv"

	"github.com/google/uuid"
)

func (ac *Service) InsertValuesToCAd(new_student *models.StudentInfoDTO, account_id string) error {

	if _, err := strconv.ParseFloat(new_student.Name, 64); err == nil {
		return fmt.Errorf("name cant be number")
	}
	course_details, err := ac.daos.GetCourseById(new_student.CourseId)
	if err != nil || course_details.Id == uuid.Nil {
		return fmt.Errorf("course Not Found")
	}

	// new_student.CourseId = course_details.Id
	sd_existing, _ := ac.daos.GetStudentDetailsByRollNumber(new_student.RollNumber)
	if sd_existing != nil {
		return fmt.Errorf("student with roll number exist")
	}

	new_student.Id = uuid.New()
	sm, err2 := ac.InsertStudentIdInToMarksTable(new_student)
	if err2 != nil {
		return err2
	}
	new_student.MarksId = sm.Id

	err1 := ac.daos.InsertIntoStudentInfos(models.StudentInfoDTOToStudentInfo(new_student))
	if err1 != nil {
		err = ac.daos.DeleteStudenetMarks(new_student.MarksId)
		if err != nil {
			log.Println("error deleting student marks")
		}
		return err1
	}
	// check if instructor is present for course
	res, err := ac.daos.GetICByInstructorIdAndCourseId(new_student.InstructorID.String(), new_student.CourseId.String())
	if err != nil || res == nil {
		err = ac.daos.DeleteStudenetMarks(new_student.MarksId)
		if err != nil {
			log.Println("error deleting student marks")
		}
		err = ac.daos.DeleteStudentDaos(new_student.Id)
		if err != nil {
			log.Println("error deleting student")
		}
		return fmt.Errorf("instructor not found for course")
	}

	err = ac.daos.InsertIntoSCI(&models.StudentCourseInstructor{
		StudentId:    new_student.Id,
		CourseId:     new_student.CourseId,
		InstructorId: new_student.InstructorID,
		Marks:        0,
		IsDeleted:    false,
	})
	if err != nil {
		return err
	}

	account := &models.Account{}
	account.Id = new_student.Id
	account.Info.Credentials.Id = new_student.Id
	account.Name = new_student.Name

	err3 := ac.daos.CreateAccount(account)
	if err3 != nil {
		log.Println("error storing in account student")
		return err3
	}
	go utils.StoreMessages("New Student", account.Name, config.AccountTypeInstructor, account_id)
	return nil

}

func (ac *Service) Retrieve_student_detailsbyOrder(order string) ([]*models.StudentInfoDTO, error) {

	rca, err := ac.daos.RetrieveCollegeAdministrationByOrder(order)

	studentsDTOS := make([]*models.StudentInfoDTO, 0)

	courses, err := ac.daos.RetieveCoursesAvailable()
	if err != nil {
		return nil, err
	}
	courseMap := make(map[uuid.UUID]*models.CourseInfo)
	for _, eachCourse := range courses {
		courseMap[eachCourse.Id] = eachCourse
	}

	for _, eachRCA := range rca {
		studentDTO := models.StudentInfoToStudentInfoDTO(eachRCA)
		courses, err := ac.daos.GetSCIByStudentId(eachRCA.Id.String())
		if err != nil {
			return nil, err
		}
		for _, eachCourse := range courses {
			courseDTO := models.StudentCourseInstructorToStudentCourseInstructorDTO(eachCourse)
			courseDTO.CourseName = courseMap[eachCourse.CourseId].CourseName
			studentDTO.ClassesEnrolled = append(studentDTO.ClassesEnrolled, courseDTO)
		}
		studentsDTOS = append(studentsDTOS, studentDTO)
	}

	for _, student := range studentsDTOS {
		for _, eachCourse := range student.ClassesEnrolled {
			marks, err := ac.daos.GetMarksByStudentIdAndCourseId(student.Id, eachCourse.CourseId)
			if err != nil {
				return nil, err
			}
			student.StudentMarks = *marks

		}
	}

	return studentsDTOS, nil
}

// Deprecated: use UpdateStudentDetailsV2 instead
func (ac *Service) Update_Student_Details(update_student *models.StudentInfoDTO) error {

	existing_student, err := ac.daos.GetStudentdetail(
		&models.StudentInfo{
			Id: update_student.Id})

	if err != nil {
		return err
	}

	if existing_student == nil {
		return fmt.Errorf("student details mismatched or does not exists")
	}
	if update_student.Age > 0 && existing_student.Age != update_student.Age {
		existing_student.Age = update_student.Age
	}
	if update_student.Name != "" && existing_student.Name != update_student.Name {
		existing_student.Name = update_student.Name
	}
	if update_student.RollNumber != "" && existing_student.RollNumber != update_student.RollNumber {
		existing_student.RollNumber = update_student.RollNumber
	}

	err5 := ac.daos.UpdateClgStudent(existing_student)
	if err5 != nil {
		return err5
	}

	return nil

}

func (ac *Service) DeleteStudent(id uuid.UUID) error {

	student, err := ac.daos.GetStudentdetail(&models.StudentInfo{Id: id})
	if err != nil {
		return err
	}
	err = ac.daos.HardDeleteStudentCourseInstructorByStudentId(student.Id.String())
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
	// TODO: REmove this function if not needed
	return nil, nil

	// course_list := make(map[string]string)
	// course_list["student_name"] = student_name
	// courses, err := ac.daos.RetieveCoursesAvailable()
	// if err != nil {
	// 	return nil, err
	// }
	// courseMap := make(map[uuid.UUID]*models.CourseInfo)
	// for _, eachCourse := range courses {
	// 	courseMap[eachCourse.Id] = eachCourse
	// }

	// si, err := ac.daos.GetSCIByStudentId(student_name)
	// if err != nil {
	// 	return nil, err
	// }
	// for index, each_si := range *si {
	// 	each_si.ClassesEnrolled, err = ac.daos.GetCourseById(each_si.CourseId)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	studentmarks, err := ac.daos.GetMarksByStudentId(each_si.Id)
	// 	each_si.StudentMarks = *studentmarks
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	course_list[fmt.Sprintf("course_%d", index+1)] = fmt.Sprintf("%s -->Marks: %d Grade: %s", each_si.ClassesEnrolled.CourseName, each_si.StudentMarks.Marks, each_si.StudentMarks.Grade)
	// }
	// return course_list, nil
}

func (ac *Service) DeleteStudentCourseService(sn, cn string) (err error) {
	// TODO: handle student course deleteion - soft delete
	// course_details, err1 := ac.daos.GetCourseByName(cn)
	// if err1 != nil {
	// 	return fmt.Errorf("course not found")
	// }
	// var student_detail models.StudentInfo
	// student_detail.Name = sn
	// student_detail.CourseId = course_details.Id

	// _, err2 := ac.daos.GetStudentdetail(&student_detail)
	// if err2 != nil {
	// 	return err2
	// }
	// err = ac.daos.DeleteCourseForAStudent(sn, course_details.Id)
	// if err != nil {
	// 	return fmt.Errorf("failed to delete")
	// }
	// err5 := ac.daos.DeleteStudenetMarks(student_detail.MarksId)
	// if err5 != nil {
	// 	return fmt.Errorf("failed to delete")
	// }
	return nil

}

func (s *Service) GetAllStudentSelectiveData() ([]*models.StudentSelectiveData, error) {

	ssd := []*models.StudentSelectiveData{}

	student_data, err := s.daos.RetrieveCollegeAdministrationByOrder("")
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
func (ac *Service) DeleteStudentSpecifics(st_req *models.StudentInfo) (err error) {

	course_details, err1 := ac.daos.GetCourseByName(st_req.ClassesEnrolled.CourseName)
	if err1 != nil {
		return fmt.Errorf("course not found")
	}

	st_details, err2 := ac.daos.GetStudentdetail(st_req)
	if err2 != nil {
		return err2
	}
	st_details.ClassesEnrolled = course_details

	err = ac.daos.DeleteStudentWithSpecifics(st_details)
	if err != nil {
		return fmt.Errorf("failed to delete")
	}
	return nil

}

func (ac *Service) UpdateStudentDetailsV2(update_student *models.StudentInfo) error {

	if _, err := strconv.ParseFloat(update_student.Name, 64); err == nil {
		return fmt.Errorf("name cant be number")
	}
	if update_student.Age <= 0 {
		return fmt.Errorf("invalid age")
	}
	existing_student, _ := ac.daos.GetStudentdetail(
		&models.StudentInfo{
			Id: update_student.Id})

	if existing_student == nil {
		return fmt.Errorf("student details mismatched or does not exists")
	}

	if update_student.Age > 0 && existing_student.Age != update_student.Age {
		existing_student.Age = update_student.Age
	}
	if update_student.Name != "" && existing_student.Name != update_student.Name {
		existing_student.Name = update_student.Name
	}
	if update_student.RollNumber != "" && existing_student.RollNumber != update_student.RollNumber {
		existing_student.RollNumber = update_student.RollNumber
	}

	err5 := ac.daos.UpdateClgStudent(existing_student)
	if err5 != nil {
		return err5
	}
	// TODO: handle courses and marks updation
	return nil

}
