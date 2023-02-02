package handlers

import (
	"CollegeAdministration/service"
	"time"

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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500", "http://localhost:5050"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE", "POST"},
		AllowHeaders:     []string{"Origin", "content-type", "Set-Cookie", "token", "account_id"},
		ExposeHeaders:    []string{"Content-Length", "Set-Cookie", "token", "account_id"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://127.0.0.1:5500" || origin == "http://localhost:5050"
		},
		MaxAge: 12 * time.Hour,
	}))

	h.RoutingChannel(&router.RouterGroup)

	return router
}
