package handlers

import (
	"CollegeAdministration/models"
	"CollegeAdministration/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func New(db *service.Service) *Handler {
	return &Handler{
		service: db,
	}
}

func (h *Handler) RoutingChannel(rc *gin.RouterGroup) {
	rc.POST("/InsertToCoursesAvailableTable", h.InsertCA)
	rc.POST("/InsertToCollegeAdministration", h.InsertCAd)
	rc.GET("/RetrieveCoursesAvailableTable", h.RetrieveValuesCA)
	rc.GET("/RetrieveCollegeAdministrationTable", h.RetrieveValuesCAd)
	rc.PATCH("/UpdateCoursesAvailableTable/:name", h.UpdateValuesCA)
	rc.PATCH("/UpdateCollegeAdministrationTable", h.UpdateValuesCAd)
	rc.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
func (h *Handler) InsertCA(ctx *gin.Context) {
	var ca models.CoursesAvailable
	err := ctx.BindJSON(&ca)

	if err != nil {
		log.Println("not able to store values")
	}
	response := h.service.InsertValuesToCA(&ca)
	if response != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error())
	} else {
		ctx.JSON(http.StatusOK, "Successfully inserted to table")
	}

}

func (h *Handler) InsertCAd(ctx *gin.Context) {
	var cad models.CollegeAdminstration
	err := ctx.BindJSON(&cad)
	if err != nil {
		log.Println("not able to store values")
	}
	response := h.service.InsertValuesToCAd(&cad)
	if response != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error())
	} else {
		ctx.JSON(http.StatusOK, "Successfully inserted to table")
	}
}

func (h *Handler) RetrieveValuesCA(ctx *gin.Context) {

	response, err := h.service.RetrieveCA()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, response)
	}
}

func (h *Handler) RetrieveValuesCAd(ctx *gin.Context) {

	response, err := h.service.RetrieveCAd()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, response)
	}
}

func (h *Handler) UpdateValuesCA(ctx *gin.Context) {
	var rc models.CoursesAvailable
	var name = ctx.Param("name")
	ctx.BindJSON(&rc)
	err := h.service.UpdateCA(name, &rc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.Error(err))
	} else {
		ctx.JSON(http.StatusOK, "Success")
	}

}

func (h *Handler) UpdateValuesCAd(ctx *gin.Context) {
	var rcd models.CollegeAdminstration
	ctx.BindJSON(&rcd)
	err := h.service.UpdateCAd(&rcd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.Error(err))
	} else {
		ctx.JSON(http.StatusOK, "Success")
	}

}
