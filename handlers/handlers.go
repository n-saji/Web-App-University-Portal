package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RoutingChannel(rc *gin.RouterGroup) {

	//course
	rc.POST("/insert-course", h.InsertCourse)
	rc.GET("/retrieve-all-courses", h.RetrieveValuesCourse)
	rc.PATCH("/update-course/:name", h.UpdateValuesCourse)
	rc.DELETE("/delete-course/:courseName", h.DeleteCourse)

	//student
	rc.POST("/insert-student-details", h.InsertStudentDetails)
	rc.GET("/retrieve-college-administration", h.RetrieveValuesForStudent)
	rc.PATCH("/update-student-details/:roll_number/:student_name/:coursename", h.UpdateValuesForStudent)
	rc.DELETE("delete-student-info/:rollnumber", h.DeleteStudentDetails)
	rc.PATCH("update-student-name-and-age/:name", h.UpdateStudentNameAndAge)
	rc.GET("find-all-course-for-student/:name", h.FetchAllCourseForAStudent)
	rc.DELETE("/delete-student-course/:name/:course", h.DeleteStudentCourse)
	rc.GET("/get-ranking/:coursename", h.GetRankingForACourse)
	rc.GET("/get-student-name-course", h.GetSelectedFieldsAllStudent)
	rc.DELETE("/delete-student", h.DeleteStudentWithSpecifics)

	//instructor
	rc.POST("/insert-instructor-details", h.AddInstructor)
	rc.GET("/retrieve-instructors", h.RetrieveInstructorDetails)
	rc.GET("/retrieve-instructors/:order_by", h.RetrieveInstructorDetailsByOrder)
	rc.GET("/instructor-login-with-id/:instructorId/:emailId/:password", h.InstructorLoginCreation)
	rc.DELETE("/delete-instructor/:name", h.DeleteInstructor)
	rc.PATCH("/update-instructor", h.UpdateInstructor)
	rc.DELETE("/delete-instructor", h.DeleteInstructorWithConditions)
	rc.GET("get-instructor-name-by-id/:id", h.GetInstructorNameWithId)
	rc.GET("/view-profile-instructor/:id", h.ViewProfile)

	//Authentication
	rc.GET("/check-token-status", h.CheckTokenStatus)
	rc.PUT("/update-instructor-credentials", h.UpdateInstructorCredentials)
	rc.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	rc.POST("/v1/login", h.login)
	rc.GET("/logout",h.logout)

	//Depreciated
	// rc.GET("/instructor-login/:emailId/:password", h.InstructorLogin)
}

/*
API TO CREATE -
sort table by each column (need to create seperate api for ordering by) - Done for students


BUG -
make buttons as icon using heroicons

Features -
re-organize the structure
MAke UI changes for each instructor by redirecting to list of students under each instructor (add a new button)
Add View PROFILE FOR INSTRUCTORS
*/
