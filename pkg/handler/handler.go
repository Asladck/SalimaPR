package handler

import (
	"SalimProject/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	h.initAuthRoutes(router)

	return router
}
func (h *Handler) initAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/refresh", h.refreshHandler)
		auth.POST("/sign-in", h.signIn)
	}
}
