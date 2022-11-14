package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) InstructorLoginCreation(ctx *gin.Context) {

	parameter := ctx.Params
	uuid, err2 := uuid.Parse(parameter.ByName("instructorId"))
	if err2 != nil {
		ctx.Abort()
		ctx.JSON(406, "invalid uuid format")
	}

	emailId := parameter.ByName("emailId")
	password := parameter.ByName("password")
	err := h.service.ValidateLogin(emailId, password)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, err.Error())
	} else if err2 == nil {
		err1 := h.service.StoreInstructoLogindetails(uuid, emailId, password)
		token, _ := h.service.GetTokenAfterLogging()
		if err1 != nil {
			ctx.JSON(http.StatusNotAcceptable, err1.Error())
		} else {
			ctx.JSON(http.StatusAccepted, fmt.Sprintf("Successfully Created Use token for accessing Db-> %s", token.String()))
		}

	}

}

func (h *Handler) InstructorLogin(ctx *gin.Context) {
	parameter := ctx.Params
	emailId := parameter.ByName("emailId")
	password := parameter.ByName("password")
	err := h.service.ValidateLogin(emailId, password)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, err.Error())
	}
	err1 := h.service.CheckCredentials(emailId, password)
	if err1 != nil {
		ctx.JSON(http.StatusUnauthorized, err1.Error())
	}

	if err == nil && err1 == nil {
		token, err2 := h.service.GetTokenAfterLogging()

		if err2 != nil {
			ctx.JSON(http.StatusNotAcceptable, err2.Error())
		} else {
			ctx.JSON(http.StatusAccepted, fmt.Sprintf("Successfully LogedIn Use token for accessing Db-> %s", token.String()))
		}

	}

}
