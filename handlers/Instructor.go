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
	URl string
}

func (h *Handler) InstructorInfoHandlers(ctx *gin.Context) {

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
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	cookie_token, err1 := ctx.Cookie("token")
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, err1.Error())
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

	rid, err := h.service.GetInstructorDetails()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusOK, rid)
	}
}

func (h *Handler) DeleteInstructor(ctx *gin.Context) {

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

	name := ctx.Param("name")
	err := h.service.DeleteInstructor(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusOK, "deleted instructor: "+name)
	}
}
