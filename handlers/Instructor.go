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

	name := ctx.Param("name")
	err := h.service.DeleteInstructor(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusOK, "deleted instructor: "+name)
	}
}
