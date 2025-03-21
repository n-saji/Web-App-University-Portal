package handlers

import (
	"CollegeAdministration/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type res struct {
	Msg string
	Err string
	URL string
}

func (h *Handler) AddInstructor(ctx *gin.Context) {

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
	var insd *models.InstructorDetailsDTO
	err := ctx.BindJSON(&insd)

	if err != nil {
		err = fmt.Errorf("unable to store values")
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	account_id, err := h.service.GetAccountByToken(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	id, response := h.service.InsertInstructor(account_id, insd)
	if response != nil {
		reply.Err = response.Error()
		ctx.JSON(http.StatusInternalServerError, &reply)
		return
	}
	reply.Msg = "successfully created. Create login for accessing db"
	reply.URL = fmt.Sprintf("/instructor-login-with-id/%s/:emailid/:password", id)
	reply.Err = "nil"

	ctx.IndentedJSON(http.StatusCreated, &reply)

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

func (h *Handler) RetrieveInstructorDetailsByOrder(ctx *gin.Context) {

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
	order_clause := ctx.Params.ByName("order_by")

	// or use order_clause =ctx.Query("order_by")

	rid, err := h.service.GetInstructorDetailsWithConditions(order_clause)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, rid)
	}
}

// deprecated function
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

	req_id := &models.InstructorDetailsDTO{}
	err := ctx.BindJSON(&req_id)
	if err != nil {
		err = fmt.Errorf("unable to store values")
		ctx.JSON(http.StatusInternalServerError, err.Error()+" err:")
		return
	}
	query := &models.InstructorDetailsDTO{}
	query_params := ctx.Query
	if val1 := query_params("instructor_code"); val1 != "" {
		query.InstructorCode = val1
	}
	if val1 := query_params("instructor_name"); val1 != "" {
		query.InstructorName = val1
	}

	if val1 := query_params("instructor_id"); val1 != "" {
		query.Id = uuid.MustParse(val1)
	}
	if query.InstructorCode == "" && query.InstructorName == "" && query.Id == uuid.Nil {
		ctx.JSON(http.StatusInternalServerError, "No query Params")
		return
	}
	err3 := h.service.Update_Instructor(req_id, *query)
	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, err3.Error())
		return
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

func (h *Handler) ViewProfile(ctx *gin.Context) {

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
	instructor_id := params.ByName("id")
	instructor_profile, err := h.service.ViewinstructorProfile(instructor_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.IndentedJSON(http.StatusOK, instructor_profile)
}

func (h *Handler) GetInstructorsForCourse(ctx *gin.Context) {
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
	course_name := params.ByName("course")
	instructor_profile, err := h.service.GetInstructorsForCourse(course_name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.IndentedJSON(http.StatusOK, instructor_profile)
}
