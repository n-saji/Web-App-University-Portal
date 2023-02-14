package handlers

import (
	"CollegeAdministration/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type res struct {
	Msg string
	Err string
	URl string
}

func (h *Handler) InstructorInfoHandlers(ctx *gin.Context) {

	token := ctx.GetHeader("Token")
	var err1 error
	if token == "" {
		token, err1 = ctx.Cookie("token")
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Sprint("no token found -", err1.Error()))
			return
		}
	}

	err2 := h.service.CheckTokenWithCookie(token)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, err2.Error())
		return
	}

	var reply res
	var insd *models.InstructorDetails
	err := ctx.BindJSON(&insd)

	if err != nil {
		err = fmt.Errorf("unable to store values")
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	id, response := h.service.InsertInstructorDet(insd)
	if response != nil {
		reply.Err = response.Error()
		ctx.JSON(http.StatusInternalServerError, &reply)
		return
	}
	reply.Msg = "successfully created. Create login for accessing db"
	reply.URl = fmt.Sprintf("/instructor-login-with-id/%s/:emailid/:password", id)
	reply.Err = "nil"

	if response == nil {
		ctx.IndentedJSON(http.StatusCreated, &reply)
	}
}

func (h *Handler) RetrieveInstructorDetails(ctx *gin.Context) {

	token := ctx.GetHeader("Token")
	var err1 error
	if token == "" {
		token, err1 = ctx.Cookie("token")
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Sprint("no token found -", err1.Error()))
			return
		}
	}

	err2 := h.service.CheckTokenWithCookie(token)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, err2.Error())
		return
	}

	rid, err := h.service.GetInstructorDetails()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, rid)
	}
}

func (h *Handler) DeleteInstructor(ctx *gin.Context) {

	token := ctx.GetHeader("Token")
	var err1 error
	if token == "" {
		token, err1 = ctx.Cookie("token")
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Sprint("no token found -", err1.Error()))
			return
		}

	}

	err2 := h.service.CheckTokenWithCookie(token)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, err2.Error())
		return
	}

	name := ctx.Param("name")
	err := h.service.DeleteInstructor(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, "deleted instructor: "+name)
	}
}

func (h *Handler) UpdateInstructor(ctx *gin.Context) {

	token := ctx.GetHeader("Token")
	var err1 error
	if token == "" {
		token, err1 = ctx.Cookie("token")
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Sprint("no token found -", err1.Error()))
			return
		}
	}

	err2 := h.service.CheckTokenWithCookie(token)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, err2.Error())
		return
	}

	req_id := &models.InstructorDetails{}
	err := ctx.BindJSON(&req_id)
	if err != nil {
		log.Println(err)
		err = fmt.Errorf("unable to store values")
		ctx.JSON(http.StatusInternalServerError, err.Error()+" err:")
		return
	}
	cond := &models.InstructorDetails{}
	query_params := ctx.Query
	if val1 := query_params("instructor_code"); val1 != "" {
		cond.InstructorCode = val1
	}
	if val1 := query_params("instructor_name"); val1 != "" {
		cond.InstructorName = val1
	}
	if val1 := query_params("course_name"); val1 != "" {
		cond.CourseName = val1
	}
	if cond.InstructorCode == "" && cond.InstructorName == "" && cond.CourseName == "" {
		ctx.JSON(http.StatusInternalServerError, "No query Params")
		return
	}
	err3 := h.service.Update_Instructor(*req_id, *cond)
	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, err3.Error())

	}
	ctx.IndentedJSON(http.StatusOK, "updated details")
}

func (h *Handler) DeleteInstructorWithConditions(ctx *gin.Context) {

	token := ctx.GetHeader("Token")
	var err1 error
	if token == "" {
		token, err1 = ctx.Cookie("token")
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Sprint("no token found -", err1.Error()))
			return
		}
	}
	err2 := h.service.CheckTokenWithCookie(token)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, err2.Error())
		return
	}

	req_id := &models.InstructorDetails{}
	err := ctx.BindJSON(&req_id)
	if err != nil {
		err = fmt.Errorf("unable to store values")
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err3 := h.service.DeleteInstructorWithConditions(req_id)
	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, err3.Error())
		return
	}
	ctx.IndentedJSON(http.StatusOK, "deleted instructor")
}

func (h *Handler) GetInstructorNameWithId(ctx *gin.Context) {
	token := ctx.GetHeader("Token")
	var err1 error
	if token == "" {
		token, err1 = ctx.Cookie("token")
		if err1 != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Sprint("no token found -", err1.Error()))
			return
		}
	}
	err2 := h.service.CheckTokenWithCookie(token)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, err2.Error())
		return
	}
	params := ctx.Params
	id := params.ByName("id")
	i_details, err := h.service.GetInstructorNamewithId(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.IndentedJSON(http.StatusOK, i_details)
}
