package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) InstructorLoginCreation(ctx *gin.Context) {

	parameter := ctx.Params
	uuid, err2 := uuid.Parse(parameter.ByName("instructorId"))
	if err2 != nil {
		ctx.JSON(406, "invalid uuid format")
		return
	}

	emailId := parameter.ByName("emailId")
	password := parameter.ByName("password")
	err := h.service.ValidateLogin(emailId, password)
	err3 := h.service.CheckEmailExist(emailId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	} else if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, err3.Error())
		return
	} else if err2 == nil {
		err1 := h.service.StoreInstructoLogindetails(uuid, emailId, password)
		if err1 != nil {
			ctx.JSON(http.StatusNotAcceptable, err1.Error())
			return
		} else {
			token, _ := h.service.GetTokenAfterLogging()
			ctx.Writer.Header().Set("token", token.String())
			ctx.Writer.Header().Set("account_id", uuid.String())
			http.SetCookie(ctx.Writer, &http.Cookie{
				Name:     "token",
				Value:    token.String(),
				Path:     "/",
				HttpOnly: true,
				Secure:   false,
			})
			http.SetCookie(ctx.Writer, &http.Cookie{
				Name:     "account_id",
				Value:    uuid.String(),
				Path:     "/",
				HttpOnly: true,
				Secure:   false,
			})
			ctx.JSON(http.StatusAccepted, "successfully created")
		}

	}

}

func (h *Handler) InstructorLogin(ctx *gin.Context) {

	parameter := ctx.Params
	emailId := parameter.ByName("emailId")
	password := parameter.ByName("password")
	err := h.service.ValidateLogin(emailId, password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err1 := h.service.CheckCredentials(emailId, password)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, err1.Error())
		return
	}
	instructor_id, err3 := h.service.GetInstructorIDWithEmail(emailId)
	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, err3.Error())
		return
	}
	if err == nil && err1 == nil {
		token, err2 := h.service.GetTokenAfterLogging()
		ctx.Writer.Header().Set("token", token.String())
		ctx.Writer.Header().Set("account_id", instructor_id)
		http.SetCookie(ctx.Writer, &http.Cookie{
			Name:     "token",
			Value:    token.String(),
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
		})
		http.SetCookie(ctx.Writer, &http.Cookie{
			Name:     "account_id",
			Value:    instructor_id,
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
		})
		if err2 != nil {
			ctx.JSON(http.StatusNotAcceptable, err2.Error())
			return
		} else {
			ctx.JSON(http.StatusAccepted, "successfully logged-in")
		}

	}

}
