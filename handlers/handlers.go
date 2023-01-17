package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RoutingChannel(rc *gin.RouterGroup) {
	rc.POST("/insert-course", h.InsertCourse)
	rc.POST("/insert-student-details", h.InsertStudentDetails)
	rc.GET("/retrieve-all-courses", h.RetrieveValuesCourse)
	rc.GET("/retrieve-college-administration", h.RetrieveValuesForStudent)
	rc.PATCH("/update-course/:name", h.UpdateValuesCourse)
	rc.PATCH("/update-student-details/:coursename", h.UpdateValuesForStudent)
	rc.DELETE("/delete-course/:courseName", h.DeleteCourse)
	rc.DELETE("delete-student-info/:rollnumber", h.DeleteStudentDetails)
	rc.PATCH("update-student-name-and-age/:name", h.UpdateStudentNameAndAge)
	rc.GET("find-all-course-for-student/:name", h.FetchAllCourseForAStudent)
	rc.POST("/insert-instructor-details", h.InstructorInfoHandlers)
	rc.GET("/retrieve-instructors", h.RetrieveInstructorDetails)
	rc.DELETE("/delete-student-course/:name/:course", h.DeleteStudentCourse)
	rc.GET("/instructor-login-with-id/:instructorId/:emailId/:password", h.InstructorLoginCreation)
	rc.GET("/instructor-login/:emailId/:password", h.InstructorLogin)
	rc.DELETE("/delete-instructor/:name", h.DeleteInstructor)
	rc.GET("/get-ranking/:coursename", h.GetRankingForACourse)
	rc.GET("/get-student-name-course", h.GetSelectedFieldsAllStudent)
	rc.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

/*
API TO CREATE -
 delete student-marks
re-organize the structure

BUG -
*/
