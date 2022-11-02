package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler)GetRouter() *gin.Engine {

	router := gin.Default()
	h.RoutingChannel(&router.RouterGroup)

	return router
}
