package handlers

import (
	"CollegeAdministration/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InstructorInfoHandlers(ctx *gin.Context) {
	var insd *models.InstructorDetails
	err := ctx.BindJSON(&insd)

	if err != nil {
		err = fmt.Errorf("unable to store values")
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := h.service.InsertInstructorDet(insd)
	if response != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error())
	} else {
		ctx.IndentedJSON(http.StatusCreated, "Successfully Inserted")
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
