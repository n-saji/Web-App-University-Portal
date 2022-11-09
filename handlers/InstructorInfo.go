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

	if err != nil{
		fmt.Errorf("unable to store values")
	}
	response := h.service.InsertInstructorDet(insd)
	if response != nil{
		ctx.JSON(http.StatusInternalServerError,response.Error())
	} else{
		ctx.IndentedJSON(http.StatusCreated,"Successfully Inserted")
	}
}
