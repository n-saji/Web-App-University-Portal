package handlers

import (
	"CollegeAdministration/models"
	"CollegeAdministration/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	token := ctx.Param("token")
	token_id, err := uuid.Parse(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("error parsing uuid").Error())
		return
	}
	status, err1 := h.service.CheckTokenValidity(token_id)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, err1.Error())
		return
	}
	if status {
		response, err := h.service.RetrieveCA()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		} else {
			ctx.JSON(http.StatusOK, response)
		}
	} else {
		ctx.JSON(http.StatusBadRequest, "token expired")
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
	type response struct {
		Message string
		Courses []models.CourseInfo
	}

	var CourseName string = ctx.Param("courseName")
	err := h.service.DeleteCA(CourseName)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		res := response{}
		token, _ := h.service.GetTokenAfterLogging()
		utils.MakeRequest(http.MethodGet, "http://localhost:5050/RetrieveCoursesAvailable/"+token.String(), "Fetching course", nil, &res.Courses)
		res.Message = "Please select from existing course"
		ctx.IndentedJSON(200, res)
	} else {
		ctx.JSON(http.StatusOK, "successfully deleted")
	}
}
