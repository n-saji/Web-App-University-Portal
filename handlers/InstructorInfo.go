package handlers

import (
	"CollegeAdministration/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) InstructorInfoHandlers(ctx *gin.Context) {
	type res struct {
		Msg string
		Err string
		URl string
	}
	var reply res
	var insd *models.InstructorDetails
	err := ctx.BindJSON(&insd)

	if err != nil {
		err = fmt.Errorf("unable to store values")
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	response, id := h.service.InsertInstructorDet(insd)
	if response == nil {
		reply.Msg = "Successfully created. Create login for accessing db"
		reply.URl = fmt.Sprintf("/instructorlogin/%s/:emailid/:password", id)
		reply.Err = "nil"
	} else {
		reply.Err = response.Error()
	}

	if response != nil {
		ctx.JSON(http.StatusInternalServerError, &reply)
	} else {
		ctx.IndentedJSON(http.StatusCreated, &reply)
	}
}

func (h *Handler) RetrieveInstructorDetails(ctx *gin.Context) {
	var rid []*models.InstructorDetails
	rid, err := h.service.GetInstructorDetails()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.IndentedJSON(http.StatusCreated, rid)
	}
}

func (h *Handler) InstructorLogin(ctx *gin.Context) {

	parameter := ctx.Params
	uuid, _ := uuid.Parse(parameter.ByName("instructorId"))
	emailId := parameter.ByName("emailId")
	password := parameter.ByName("password")
	err := h.service.ValidateLogin(emailId, password)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, err.Error())
	} else {
		err1 := h.service.StoreInstructoLogindetails(uuid, emailId, password)
		if err1 != nil {
			ctx.JSON(http.StatusNotAcceptable, err1.Error())
		} else {
			ctx.JSON(http.StatusAccepted, "Successfully Created Use token for accessing Db ")
		}

	}

}
