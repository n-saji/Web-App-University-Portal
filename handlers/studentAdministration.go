package handlers

import (
	"CollegeAdministration/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InsertCAd(ctx *gin.Context) {
	var cad models.CollegeAdminstration
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

	response, err := h.service.RetrieveCAd()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, response)
	}
}



func (h *Handler) UpdateValuesCAd(ctx *gin.Context) {
	var rcd models.CollegeAdminstration
	ctx.BindJSON(&rcd)
	err := h.service.UpdateCAd(&rcd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.Error(err))
	} else {
		ctx.JSON(http.StatusOK, "Success")
	}

}

