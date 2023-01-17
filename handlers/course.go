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

func (h *Handler) InsertCourse(ctx *gin.Context) {

	token, err3 := ctx.Cookie("token")
	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, err3.Error())
		return
	}
	token_id, err4 := uuid.Parse(token)
	if err4 != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("error parsing uuid").Error())
		return
	}
	status, err1 := h.service.CheckTokenValidity(token_id)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, err1.Error())
		return
	}
	if !status {
		ctx.JSON(http.StatusBadRequest, "token expired")
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

	//ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ok := ctx.GetHeader("Internal-call")

	if ok != "true" {
		cookie_token, err3 := ctx.Cookie("token")
		if err3 != nil {
			ctx.JSON(http.StatusInternalServerError, err3.Error())
			return
		}

		if cookie_token == "" {
			ctx.JSON(http.StatusInternalServerError, "authentication error")
			return
		}

		validity, _ := h.service.CheckTokenValidity(uuid.MustParse(cookie_token))
		if !validity {
			ctx.JSON(http.StatusInternalServerError, "authentication time-out")
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

	token, err3 := ctx.Cookie("token")
	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, err3.Error())
		return
	}
	token_id, err4 := uuid.Parse(token)
	if err4 != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("error parsing uuid").Error())
		return
	}
	status, err1 := h.service.CheckTokenValidity(token_id)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, err1.Error())
		return
	}
	if !status {
		ctx.JSON(http.StatusBadRequest, "token expired")
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

	token, err3 := ctx.Cookie("token")
	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, err3.Error())
		return
	}
	token_id, err4 := uuid.Parse(token)
	if err4 != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("error parsing uuid").Error())
		return
	}
	status, err1 := h.service.CheckTokenValidity(token_id)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, err1.Error())
		return
	}
	if !status {
		ctx.JSON(http.StatusBadRequest, "token expired")
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
