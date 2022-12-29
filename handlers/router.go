package handlers

import (
	"CollegeAdministration/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func New(db *service.Service) *Handler {
	return &Handler{
		service: db,
	}
}

func (h *Handler) GetRouter() *gin.Engine {

	router := gin.Default()
	router.Use(cors.Default())
	h.RoutingChannel(&router.RouterGroup)

	return router
}
