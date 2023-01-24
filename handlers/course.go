package handlers

import (
	"CollegeAdministration/models"
	"CollegeAdministration/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InsertCourse(ctx *gin.Context) {

	token := ctx.GetHeader("Token")
	var err1 error
	if token == "" {
		token, err1 = ctx.Cookie("token")
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, err1.Error())
			return
		}
	}

	err2 := h.service.CheckTokenWithCookie(token)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, err2.Error())
		return
	}

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

func (h *Handler) RetrieveValuesCourse(ctx *gin.Context) {

	ok := ctx.GetHeader("Internal-call")

	if ok != "true" {
		token := ctx.GetHeader("Token")
		var err1 error
		if token == "" {
			token, err1 = ctx.Cookie("token")
			if err1 != nil {
				ctx.JSON(http.StatusInternalServerError, err1.Error())
				return
			}
		}

		err2 := h.service.CheckTokenWithCookie(token)
		if err2 != nil {
			ctx.JSON(http.StatusInternalServerError, err2.Error())
			return
		}
	}

	response, err := h.service.RetrieveCA()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, response)
	}

}
func (h *Handler) UpdateValuesCourse(ctx *gin.Context) {

	token := ctx.GetHeader("Token")
	var err1 error
	if token == "" {
		token, err1 = ctx.Cookie("token")
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, err1.Error())
			return
		}
	}

	err2 := h.service.CheckTokenWithCookie(token)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, err2.Error())
		return
	}
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

func (h *Handler) DeleteCourse(ctx *gin.Context) {

	token := ctx.GetHeader("Token")
	var err1 error
	if token == "" {
		token, err1 = ctx.Cookie("token")
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, err1.Error())
			return
		}
	}

	err2 := h.service.CheckTokenWithCookie(token)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, err2.Error())
		return
	}

	var CourseName string = ctx.Param("courseName")
	err := h.service.DeleteCA(CourseName)
	if err != nil {
		res := models.DeleteResponse{}
		_, err1 := utils.MakeRequest(http.MethodGet, "http://localhost:5050/retrieve-all-courses", "Fetching course", nil, &res.Courses)
		if err1 != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, err1)
			return
		}
		res.Message = err.Error() + "! Please select from existing course"
		ctx.IndentedJSON(http.StatusNotFound, res)
	} else {
		ctx.JSON(http.StatusOK, "successfully deleted")
	}
}
