package models

import "github.com/google/uuid"

type InstructorDetailsDTO struct {
	Id             uuid.UUID        `json:"id"`
	InstructorCode string           `json:"instructor_code"`
	InstructorName string           `json:"instructor_name"`
	Department     string           `json:"department"`
	CourseId       uuid.UUID        `json:"course_id"`
	Courses        []*CourseInfoDTO `json:"courses"`
	StudentsLimit  int64            `json:"students_limit"`
	TotalCourses   int64            `json:"total_courses"`
	TotalStudents  int64            `json:"total_students"`
}

type InstructorProfile struct {
	Name        string           `json:"name"`
	CourseList  string           `json:"course_list"`
	Department  string           `json:"department"`
	Code        string           `json:"code"`
	Courses     []*CourseInfoDTO `json:"courses"`
	Credentials InstructorLogin  `json:"credentials"`
}

type CourseInfoDTO struct {
	Id               uuid.UUID `json:"id"`
	CourseName       string    `json:"course_name"`
	StudentsLimit    int64     `json:"students_limit"`
	StudentsEnrolled int64     `json:"students_enrolled"`
	CourseRating     int64     `json:"course_rating"`
	IsDeleted        bool      `json:"is_deleted"`
}

type StudentInfoDTO struct {
	Id              uuid.UUID                     `json:"id"`
	Name            string                        `json:"name"`
	RollNumber      string                        `json:"roll_number"`
	Age             int64                         `json:"age"`
	CourseId        uuid.UUID                     `json:"course_id"`
	MarksId         uuid.UUID                     `json:"marks_id"`
	InstructorID    uuid.UUID                     `json:"instructor_id"`
	ClassesEnrolled []*StudentCourseInstructorDTO `json:"classes_enrolled"`
	StudentMarks    StudentMarks                  `json:"student_marks"`
}

type StudentCourseInstructorDTO struct {
	StudentId    uuid.UUID `json:"student_id"`
	CourseId     uuid.UUID `json:"course_id"`
	InstructorId uuid.UUID `json:"instructor_id"`
	CourseName   string    `json:"course_name"`
	Marks        int64     `json:"marks"`
	IsDeleted    bool      `json:"is_deleted"`
}

func StudentCourseInstructorDTOToStudentCourseInstructor(studentCourseInstructorDTO *StudentCourseInstructorDTO) *StudentCourseInstructor {
	return &StudentCourseInstructor{
		StudentId:    studentCourseInstructorDTO.StudentId,
		CourseId:     studentCourseInstructorDTO.CourseId,
		InstructorId: studentCourseInstructorDTO.InstructorId,
		Marks:        studentCourseInstructorDTO.Marks,
		IsDeleted:    studentCourseInstructorDTO.IsDeleted,
	}
}

func StudentCourseInstructorToStudentCourseInstructorDTO(studentCourseInstructor *StudentCourseInstructor) *StudentCourseInstructorDTO {
	return &StudentCourseInstructorDTO{
		StudentId:    studentCourseInstructor.StudentId,
		CourseId:     studentCourseInstructor.CourseId,
		InstructorId: studentCourseInstructor.InstructorId,
		Marks:        studentCourseInstructor.Marks,
		IsDeleted:    studentCourseInstructor.IsDeleted,
	}
}

func StudentInfoDTOToStudentInfo(studentInfoDTO *StudentInfoDTO) *StudentInfo {
	return &StudentInfo{
		Id:           studentInfoDTO.Id,
		Name:         studentInfoDTO.Name,
		RollNumber:   studentInfoDTO.RollNumber,
		Age:          studentInfoDTO.Age,
		StudentMarks: studentInfoDTO.StudentMarks,
	}
}

func StudentInfoToStudentInfoDTO(studentInfo *StudentInfo) *StudentInfoDTO {
	return &StudentInfoDTO{
		Id:           studentInfo.Id,
		Name:         studentInfo.Name,
		RollNumber:   studentInfo.RollNumber,
		Age:          studentInfo.Age,
		StudentMarks: studentInfo.StudentMarks,
	}
}

func InstructorDetailsDTOToInstructorDetails(instructorDetailsDTO *InstructorDetailsDTO) *InstructorDetails {
	return &InstructorDetails{
		Id:             instructorDetailsDTO.Id,
		InstructorCode: instructorDetailsDTO.InstructorCode,
		InstructorName: instructorDetailsDTO.InstructorName,
		Department:     instructorDetailsDTO.Department,
	}
}

func InstructorDetailsToInstructorDetailsDTO(instructorDetails *InstructorDetails) *InstructorDetailsDTO {
	return &InstructorDetailsDTO{
		Id:             instructorDetails.Id,
		InstructorCode: instructorDetails.InstructorCode,
		InstructorName: instructorDetails.InstructorName,
		Department:     instructorDetails.Department,
	}
}
