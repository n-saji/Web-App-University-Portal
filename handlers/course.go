package handlers

import (
	"CollegeAdministration/models"
	"CollegeAdministration/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InsertCA(ctx *gin.Context) {
	var ca models.CourseInfo
	err := ctx.BindJSON(&ca)

	if err != nil {
		log.Println("not able to store values")
	}
	response := h.service.InsertValuesToCA(&ca)
	if response != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error())
	} else {
		ctx.JSON(http.StatusCreated, "successfully inserted to table")
	}

}

func (h *Handler) RetrieveValuesCA(ctx *gin.Context) {

	response, err := h.service.RetrieveCA()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, response)
	}

}
func (h *Handler) UpdateValuesCA(ctx *gin.Context) {
	var rc models.CourseInfo
	var name = ctx.Param("name")
	ctx.BindJSON(&rc)
	err := h.service.UpdateCA(name, &rc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.Error(err))
	} else {
		ctx.JSON(http.StatusOK, "successfully updated")
	}

}

func (h *Handler) DeleteCA(ctx *gin.Context) {

	var CourseName string = ctx.Param("courseName")
	err := h.service.DeleteCA(CourseName)
	if err != nil {
		res := models.DeleteResponse{}
		token, _ := h.service.GetTokenAfterLogging()
		utils.MakeRequest(http.MethodGet, "http://localhost:5050/retrieve-all-courses/"+token.String(), "Fetching course", nil, &res.Courses)
		res.Message = err.Error() + "! Please select from existing course"
		ctx.IndentedJSON(http.StatusNotFound, res)
	} else {
		ctx.JSON(http.StatusOK, "successfully deleted")
	}
}
