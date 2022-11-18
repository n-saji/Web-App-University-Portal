package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RoutingChannel(rc *gin.RouterGroup) {
	rc.POST("/InsertToCoursesAvailable", h.InsertCA)
	rc.POST("/InsertToCollegeAdministration", h.InsertCAd)
	rc.GET("/RetrieveCoursesAvailable/:token", h.RetrieveValuesCA)
	rc.GET("/RetrieveCollegeAdministration/:token", h.RetrieveValuesCAd)
	rc.PATCH("/UpdateCoursesAvailable/:name", h.UpdateValuesCA)
	rc.PATCH("/UpdateCollegeAdministration/:coursename", h.UpdateValuesCAd)
	rc.DELETE("/DeleteCoursesAvailable/:courseName", h.DeleteCA)
	rc.DELETE("DeleteStudentInfo/:rollnumber", h.DeleteSA)
	rc.PATCH("UpdateStudentNameAndAge/:name", h.UpdateStudentNameAndAge)
	rc.GET("FindAllCourseForAStudent/:token/:name", h.FetchAllCourseForAStudent)
	rc.POST("/InsertInstructorDetails", h.InstructorInfoHandlers)
	rc.GET("/RetrieveInstructors", h.RetrieveInstructorDetails)
	rc.DELETE("/DeleteStudentCourse/:name/:course", h.DeleteStudentCourse)
	rc.GET("/instructorlogin/:instructorId/:emailId/:password", h.InstructorLoginCreation)
	rc.GET("/instructor-login/:emailId/:password", h.InstructorLogin)
	rc.DELETE("/delete-instructor/:name", h.DeleteInstructor)
	rc.GET("/get-ranking/:token/:coursename", h.GetRankingForACourse)
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
*/
