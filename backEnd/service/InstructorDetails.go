package service

import (
	"CollegeAdministration/config"
	"CollegeAdministration/models"
	"CollegeAdministration/utils"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (ac *Service) InsertInstructor(account_id string, iid *models.InstructorDetailsDTO) (uuid.UUID, error) {
	if _, err := strconv.Atoi(iid.InstructorName); err == nil {
		return uuid.Nil, errors.New("name can't be number")
	}
	if _, err := strconv.Atoi(iid.Department); err == nil {
		return uuid.Nil, errors.New("departmment can't be number")
	}
	cn, err1 := ac.daos.GetCourseById(iid.CourseId)
	if err1 != nil || cn.Id == uuid.Nil {
		return uuid.Nil, fmt.Errorf("course not available")
	}

	cd_exist, _ := ac.daos.GetInstructor(
		models.InstructorDetailsDTOToInstructorDetails(&models.InstructorDetailsDTO{
			InstructorCode: iid.InstructorCode,
		}))
	if cd_exist.Id != uuid.Nil {
		return uuid.Nil, fmt.Errorf("instructor exits")
	}
	ok, err2 := ac.ValidateInstructorDetails(iid)
	if !ok || err2 != nil {
		return uuid.Nil, err2
	}

	iid.Id = uuid.New()

	err := ac.daos.InsertInstructorDetails(models.InstructorDetailsDTOToInstructorDetails(iid))
	if err != nil {
		log.Println("Error while inserting details")
		return uuid.Nil, err
	}

	err = ac.daos.InsertIntoIC(&models.InstructorCourse{
		InstructorId:     iid.Id,
		CourseId:         iid.CourseId,
		StudentsLimit:    iid.StudentsLimit,
		StudentsEnrolled: 0,
		CourseRating:     0,
		IsDeleted:        false,
	})
	if err != nil {
		log.Println("Error while inserting into instructor course")
		return uuid.Nil, err
	}

	account := &models.Account{}
	account.Id = iid.Id
	account.Info.Credentials.Id = iid.Id
	account.Name = iid.InstructorName
	account.Type = config.AccountTypeInstructor

	err3 := ac.daos.CreateAccount(account)
	if err3 != nil {
		log.Println("error storing in account")
		return uuid.Nil, nil
	}

	go utils.StoreMessages("New Instructor", account.Name, config.AccountTypeInstructor, account_id)

	return iid.Id, nil
}

func (ac *Service) GetInstructorDetails() ([]*models.InstructorDetailsDTO, error) {

	id, err := ac.daos.GetAllInstructor()
	if err != nil {
		return nil, err
	}

	courses, err := ac.daos.RetieveCoursesAvailable()
	if err != nil {
		return nil, err
	}
	courseMap := make(map[uuid.UUID]*models.CourseInfo)
	for _, eachCourse := range courses {
		courseMap[eachCourse.Id] = eachCourse
	}

	InstructorsDTOS := make([]*models.InstructorDetailsDTO, 0)
	for _, eachId := range id {
		InstructorsDTOS = append(InstructorsDTOS, models.InstructorDetailsToInstructorDetailsDTO(eachId))
	}
	for _, instr := range InstructorsDTOS {

		iCourses, err := ac.daos.GetCoursesByInstructorId(instr.Id.String())
		if err != nil {
			return nil, err
		}
		for _, eachCourse := range iCourses {
			courses := &models.CourseInfoDTO{
				Id:               eachCourse.CourseId,
				CourseName:       courseMap[eachCourse.CourseId].CourseName,
				StudentsLimit:    eachCourse.StudentsLimit,
				StudentsEnrolled: eachCourse.StudentsEnrolled,
				CourseRating:     eachCourse.CourseRating,
				IsDeleted:        eachCourse.IsDeleted,
			}
			instr.Courses = append(instr.Courses, courses)
		}
		res, err := ac.daos.GetSCIByInstructorId(instr.Id.String())
		if err != nil {
			return nil, err
		}
		instr.TotalStudents = int64(len(res))
		instr.TotalCourses = int64(len(iCourses))
	}

	return InstructorsDTOS, nil
}

func (s *Service) GetInstructorDetailsWithConditions(order_clause string) ([]*models.InstructorDetailsDTO, error) {

	instructors, err := s.daos.GetAllInstructorOrderByCondition(order_clause)
	if err != nil {
		return nil, err
	}
	InstructorsDTOS := make([]*models.InstructorDetailsDTO, 0)
	for _, eachId := range instructors {
		InstructorsDTOS = append(InstructorsDTOS, models.InstructorDetailsToInstructorDetailsDTO(eachId))
	}
	courses, err := s.daos.RetieveCoursesAvailable()
	if err != nil {
		return nil, err
	}
	courseMap := make(map[uuid.UUID]*models.CourseInfo)
	for _, eachCourse := range courses {
		courseMap[eachCourse.Id] = eachCourse
	}
	for _, instr := range InstructorsDTOS {

		iCourses, err := s.daos.GetCoursesByInstructorId(instr.Id.String())
		if err != nil {
			return nil, err
		}
		for _, eachCourse := range iCourses {
			courses := &models.CourseInfoDTO{
				Id:               eachCourse.CourseId,
				CourseName:       courseMap[eachCourse.CourseId].CourseName,
				StudentsLimit:    eachCourse.StudentsLimit,
				StudentsEnrolled: eachCourse.StudentsEnrolled,
				CourseRating:     eachCourse.CourseRating,
				IsDeleted:        eachCourse.IsDeleted,
			}
			instr.Courses = append(instr.Courses, courses)
		}
		res, err := s.daos.GetSCIByInstructorId(instr.Id.String())
		if err != nil {
			return nil, err
		}
		instr.TotalStudents = int64(len(res))
		instr.TotalCourses = int64(len(iCourses))
	}

	return InstructorsDTOS, nil
}

func (ac *Service) StoreInstructoLogindetails(id uuid.UUID, emailid, password string) error {

	var credentials models.InstructorLogin

	crypted_password, err2 := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err2 != nil {
		return fmt.Errorf("error parsing password")
	}
	credentials.Id = id
	credentials.EmailId = emailid
	credentials.Password = string(crypted_password)

	if id == uuid.Nil {
		return fmt.Errorf("uuid cant be null")
	}
	err := ac.daos.CheckIDPresent(credentials.Id)
	if err != nil {
		return err
	}
	err = ac.daos.CreateInstructorLogin(credentials)
	if err != nil {
		return err
	}
	iid, err := ac.daos.GetInstructor(&models.InstructorDetails{Id: id})
	if err != nil {
		return err
	}
	account := &models.Account{}
	account.Id = iid.Id
	account.Info.Credentials = models.InstructorLogin{Id: iid.Id,
		EmailId:  emailid,
		Password: string(crypted_password)}
	account.Name = iid.InstructorName

	err3 := ac.daos.AccountMigrationsUpdate([]*models.Account{account})
	if err3 != nil {
		log.Println("error storing in account instructor")
		return err3
	}
	return nil
}

func (s *Service) DeleteInstructor(name string) error {

	id, err := s.daos.GetInstructorWithName(name)
	if id.Id == uuid.Nil {
		return fmt.Errorf("instructor not found")
	}
	if err != nil {
		return err
	}

	err3 := s.daos.DeleteMessageByAccountId(id.Id)
	if err3 != nil {
		return err3
	}

	err1 := s.daos.DeleteInstructorLogin(id.Id)
	if err1 != nil {
		return err1
	}

	err2 := s.daos.DeleteInstructor(name)
	if err2 != nil {
		return err2
	}
	return nil
}

func (s *Service) Update_Instructor(req_id *models.InstructorDetailsDTO, cond models.InstructorDetailsDTO) error {

	ok, err2 := s.ValidateInstructorDetails(req_id)
	if !ok {
		return err2
	}

	list_details, err := s.GetInstructorDetailWithSpecifics(*models.InstructorDetailsDTOToInstructorDetails(&cond))
	if err != nil {
		return fmt.Errorf(fmt.Sprint("fetching error " + err.Error()))
	}
	if len(list_details) == 0 {
		return fmt.Errorf("no instructor found with given details")
	}
	old_instructor, err := s.daos.GetInstructor(&models.InstructorDetails{Id: cond.Id})
	if err != nil {
		return fmt.Errorf("no instructor found with given details")
	}
	if req_id.InstructorName != "" {
		old_instructor.InstructorName = req_id.InstructorName
	}
	if req_id.Department != "" {
		old_instructor.Department = req_id.Department
	}
	if req_id.InstructorCode != "" {
		old_instructor.Department = req_id.Department
	}
	if req_id.Courses != nil {
		for _, eachCourse := range req_id.Courses {
			_, err := s.daos.GetCourseById(eachCourse.Id)
			if err != nil {
				return fmt.Errorf("course not found")
			}
			ICDao, err := s.daos.GetICByInstructorIdAndCourseId(old_instructor.Id.String(), eachCourse.Id.String())
			if err != nil {
				return fmt.Errorf("unable to find any details from instructor-course table")
			}
			ICDao.StudentsLimit = eachCourse.StudentsLimit
			ICDao.StudentsEnrolled = eachCourse.StudentsEnrolled
			ICDao.CourseRating = eachCourse.CourseRating
			ICDao.IsDeleted = eachCourse.IsDeleted
			err = s.daos.UpdateInstructorCourse(ICDao)
			if err != nil {
				return fmt.Errorf("course not found")
			}
		}
	}

	err = s.daos.UpdateInstructor(models.InstructorDetailsDTOToInstructorDetails(req_id), models.InstructorDetailsDTOToInstructorDetails(&cond))
	if err != nil {
		return fmt.Errorf("error updating instructor details", err)
	}

	return nil

}

func (s *Service) GetInstructorDetailWithSpecifics(req models.InstructorDetails) ([]*models.InstructorDetails, error) {

	id_list, err := s.daos.RetieveInstructorDetailsWithCondition(req)
	if err != nil {
		return nil, fmt.Errorf("error %s", err.Error())
	}

	return id_list, nil

}

func (s *Service) DeleteInstructorWithConditions(id_condition *models.InstructorDetails) error {

	id_list, err := s.daos.RetieveInstructorDetailsWithCondition(*id_condition)
	if err != nil {
		return err
	}
	fmt.Println(id_list)
	for _, each_id := range id_list {
		if each_id.Id == uuid.Nil {
			return fmt.Errorf("instructor not found")
		}

		coursesLinked, err := s.daos.GetAllCoursesByInstructorId(each_id.Id.String())
		if err != nil {
			return err
		}
		if len(coursesLinked) > 0 {
			for _, eachCourse := range coursesLinked {
				err = s.daos.HardDeleteInstructorCourse(each_id.Id.String(), eachCourse.CourseId.String())
				if err != nil {
					return err
				}
			}
		}

		err3 := s.daos.DeleteMessageByAccountId(each_id.Id)
		if err3 != nil {
			return err3
		}

		err1 := s.daos.DeleteInstructorLogin(each_id.Id)
		if err1 != nil {
			return err1
		}

		err2 := s.daos.DeleteInstructorWithConditions(each_id)
		if err2 != nil {
			return err2
		}
	}

	return nil
}

func (s *Service) GetInstructorIDWithEmail(email string) (string, error) {

	instructor_id, err := s.daos.GetIDUsingEmail(email)
	if err != nil {
		return "", fmt.Errorf("failed fetching :" + err.Error())
	}
	return instructor_id, nil
}

func (s *Service) GetInstructorNamewithId(id string) (*models.InstructorDetails, error) {
	iid := &models.InstructorDetails{}
	id_uuid, err1 := uuid.Parse(id)
	if err1 != nil {
		return nil, err1
	}
	iid.Id = id_uuid
	i_details, err := s.daos.GetInstructor(iid)
	if err != nil {
		return nil, err
	}
	return &models.InstructorDetails{
		InstructorName: i_details.InstructorName,
	}, nil
}

func (s *Service) ViewinstructorProfile(i_id string) (*models.InstructorProfile, error) {
	Profile := &models.InstructorProfile{}
	i_id_parsed, err := uuid.Parse(i_id)
	if err != nil {
		return nil, err
	}
	instructor_detail, err := s.daos.GetInstructor(&models.InstructorDetails{Id: i_id_parsed})
	if err != nil {
		return nil, err
	}
	credentials, err1 := s.daos.FetchCredentialsUsingID(i_id_parsed)
	if err1 != nil {
		return nil, err1
	}
	courses, err := s.daos.GetCoursesByInstructorId(i_id)
	if err != nil {
		return nil, err
	}
	Profile.Courses = make([]*models.CourseInfoDTO, 0)
	for _, eachCourse := range courses {
		course, err := s.daos.GetCourseById(eachCourse.CourseId)
		if err != nil {
			return nil, err
		}
		Icourse, err := s.daos.GetICByInstructorIdAndCourseId(i_id, eachCourse.CourseId.String())
		if err != nil {
			return nil, err
		}
		courseInfo := &models.CourseInfoDTO{
			Id:               course.Id,
			CourseName:       course.CourseName,
			StudentsLimit:    Icourse.StudentsLimit,
			StudentsEnrolled: Icourse.StudentsEnrolled,
			CourseRating:     Icourse.CourseRating,
			IsDeleted:        Icourse.IsDeleted,
		}
		Profile.Courses = append(Profile.Courses, courseInfo)
	}
	Profile.Code = instructor_detail.InstructorCode
	Profile.Name = instructor_detail.InstructorName
	Profile.Department = instructor_detail.Department
	Profile.Credentials = models.InstructorLogin{Id: credentials.Id,
		EmailId:  credentials.EmailId,
		Password: credentials.Password}

	return Profile, nil
}

func (s *Service) UpdateInstructorCredentials(cred *models.InstructorLogin) error {

	var crypted_password []byte
	var err1 error

	if cred.EmailId == "" {
		existing_credentials, err := s.daos.FetchCredentialsUsingID(cred.Id)
		if err != nil {
			return err
		}
		cred.EmailId = existing_credentials.EmailId
	} else {
		err2 := s.ValidateEmail(cred.EmailId)
		if err2 != nil {
			return err2
		}
		err := s.CheckEmailExist(cred.EmailId)
		if err != nil {
			return err
		}
	}

	if cred.Password != "" {
		err2 := s.ValidatePassword(cred.Password)
		if err2 != nil {
			return err2
		}
		crypted_password, err1 = bcrypt.GenerateFromPassword([]byte(cred.Password), 10)
		if err1 != nil {
			return err1
		}
		cred.Password = string(crypted_password)
	} else {
		existing_credentials, err := s.daos.FetchCredentialsUsingID(cred.Id)
		if err != nil {
			return err
		}
		cred.Password = existing_credentials.Password
	}

	err := s.daos.UpdateCredentials(cred)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetInstructorsForCourse(courseName string) ([]*models.InstructorDetailsDTO, error) {

	course, err := s.daos.GetCourseByName(courseName)
	if err != nil {
		return nil, err
	}

	instructors, err := s.daos.GetInstructorsByCourseId(course.Id.String())
	if err != nil {
		return nil, err
	}

	InstructorsDTOS := make([]*models.InstructorDetailsDTO, 0)
	for _, eachId := range instructors {
		if eachId.StudentsEnrolled >= eachId.StudentsLimit {
			continue
		}
		inst_name, err := s.GetInstructorNamewithId(eachId.InstructorId.String())
		if err != nil {
			return nil, err
		}
		InstructorsDTOS = append(InstructorsDTOS, &models.InstructorDetailsDTO{
			Id:             eachId.InstructorId,
			InstructorName: inst_name.InstructorName,
		})
	}

	return InstructorsDTOS, nil
}
