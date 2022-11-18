package handlers

import (
	"CollegeAdministration/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) InsertCAd(ctx *gin.Context) {
	var cad models.StudentInfo
	err := ctx.BindJSON(&cad)
	if err != nil {
		log.Println("not able to store values")
	}
	response := h.service.InsertValuesToCAd(&cad)
	if response != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error())
	} else {
		ctx.JSON(http.StatusCreated, "Successfully inserted to table")
	}
}

func (h *Handler) RetrieveValuesCAd(ctx *gin.Context) {

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
		response, err := h.service.RetrieveCAd()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		} else {
			ctx.JSON(http.StatusOK, response)
		}
	} else {
		ctx.JSON(http.StatusBadRequest, "Token Expired")
	}
}

func (h *Handler) UpdateValuesCAd(ctx *gin.Context) {

	oldCourse := ctx.Param("coursename")
	var rcd models.StudentInfo
	ctx.BindJSON(&rcd)
	err := h.service.UpdateCAd(&rcd, oldCourse)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.Error(err))
	} else {
		ctx.JSON(http.StatusOK, "Success")
	}

}

func (h *Handler) DeleteSA(ctx *gin.Context) {

	rollNumber := ctx.Param("rollnumber")

	err := h.service.DeleteStudent(rollNumber)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.Error(err))
	} else {
		ctx.JSON(http.StatusOK, "Success")
	}

}
func (h *Handler) UpdateStudentNameAndAge(ctx *gin.Context) {
	var rcd *models.StudentInfo
	existing_name := ctx.Param("name")
	ctx.BindJSON(&rcd)
	err := h.service.UpdateStudentNameAge(existing_name, rcd.Name, rcd.Age)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.Error(err))
	} else {
		ctx.JSON(http.StatusOK, "Success")
	}

}

func (h *Handler) FetchAllCourseForAStudent(ctx *gin.Context) {

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
		student_name := ctx.Param("name")
		res, err := h.service.FetchStudentCourse(student_name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, error.Error(err))
		} else {
			ctx.JSON(http.StatusOK, res)
		}
	} else {
		ctx.JSON(http.StatusBadRequest, "Token Expired")
	}

}

func (h *Handler) DeleteStudentCourse(ctx *gin.Context) {

	parameter := ctx.Params
	student_name := parameter.ByName("name")
	course_name := parameter.ByName("course")
	err := h.service.DeleteStudentCourseService(student_name, course_name)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusOK, "Deleted")
	}
}

func (h *Handler) GetRankingForACourse(ctx *gin.Context) {

	token := ctx.Param("token")
	course_name := ctx.Param("coursename")
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
		model, err := h.service.GetAllStudentsMarksForGivenCourse(course_name)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
		} else {
			ctx.IndentedJSON(http.StatusOK, model)
		}
	} else {
		ctx.JSON(http.StatusBadRequest, "Token Expired")
	}

}
