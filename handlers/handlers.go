package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RoutingChannel(rc *gin.RouterGroup) {
	rc.POST("/insert-course", h.InsertCA)
	rc.POST("/insert-student-details", h.InsertCAd)
	rc.GET("/retrieve-all-courses", h.RetrieveValuesCA)
	rc.GET("/retrieve-college-administration", h.RetrieveValuesCAd)
	rc.PATCH("/update-course/:name", h.UpdateValuesCA)
	rc.PATCH("/update-student-details/:coursename", h.UpdateValuesCAd)
	rc.DELETE("/delete-course/:courseName", h.DeleteCA)
	rc.DELETE("delete-student-info/:rollnumber", h.DeleteSA)
	rc.PATCH("update-student-same-and-age/:name", h.UpdateStudentNameAndAge)
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
