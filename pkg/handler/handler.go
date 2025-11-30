package handler

import (
	"SalimProject/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
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
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
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

func (h *Handler) initApiRoutes(router *gin.Engine) {
	api := router.Group("/api", h.userIdentity)
	{
		clothes := api.Group("/clothes")
		{
			clothes.POST("/item", h.addClothes)
			clothes.GET("/item/:id", h.getClothesByUserId)
			clothes.DELETE("/item/:id", h.deleteClothes)
			clothes.PUT("/item/:id", h.updateClothes)
		}
	}
}
