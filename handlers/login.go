package handlers

import (
	"CollegeAdministration/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) login(ctx *gin.Context) {

	req := &models.InstructorLogin{}

	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusMethodNotAllowed, err)
		return
	}
	err = h.service.ValidateLogin(req.EmailId, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err1 := h.service.CheckCredentials(req.EmailId, req.Password)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, err1.Error())
		return
	}
	instructor_id, err3 := h.service.GetInstructorIDWithEmail(req.EmailId)
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
			ctx.JSON(http.StatusOK, "successfully logged-in")
		}

	}
}
