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

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Serve static files
	router.Static("/css", "./public/css")
	router.Static("/js", "./public/js")
	router.StaticFile("/", "./public/index.html")
	router.StaticFile("/index.html", "./public/index.html")
	router.StaticFile("/login.html", "./public/login.html")
	router.StaticFile("/signup.html", "./public/signup.html")
	router.StaticFile("/wardrobe.html", "./public/wardrobe.html")
	router.StaticFile("/outfit.html", "./public/outfit.html")

	h.initAuthRoutes(router)
	h.initApiRoutes(router)
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
			clothes.GET("/item", h.getAllClothes)
			clothes.GET("/item/:id", h.getClothesById)
			clothes.DELETE("/item/:id", h.deleteClothesById)
			clothes.PUT("/item/:id", h.updateClothesById)
		}
		outfit := api.Group("/outfit")
		{
			outfit.POST("/generate", h.generateOutfit)
			outfit.GET("/", h.getAllOutfits)
			outfit.GET("/:id", h.getOutfitById)
			outfit.DELETE("/:id", h.deleteOutfitById)
		}

	}
}
