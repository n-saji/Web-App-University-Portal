package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RoutingChannel(rc *gin.RouterGroup) {
	rc.POST("/InsertToCoursesAvailable", h.InsertCA)
	rc.POST("/InsertToCollegeAdministration", h.InsertCAd)
	rc.GET("/RetrieveCoursesAvailable", h.RetrieveValuesCA)
	rc.GET("/RetrieveCollegeAdministration", h.RetrieveValuesCAd)
	rc.PATCH("/UpdateCoursesAvailable/:name", h.UpdateValuesCA)
	rc.PATCH("/UpdateCollegeAdministration", h.UpdateValuesCAd)
	rc.DELETE("/DeleteCooursesAvailable/:courseName", h.DeleteCA)
	rc.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
