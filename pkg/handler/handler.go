package handler

import (
	"github.com/gin-gonic/gin"
	"ipk/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}
	api := router.Group("/api", h.userIdentity)
	{
		chair := api.Group("/chair")
		{
			chair.GET("/getall", h.getAllChairs)
			chair.POST("/create", h.createChair)
		}
		stat := api.Group("/stat")
		{
			stat.GET("/")
		}

		test := api.Group("/test")
		{
			test.GET("/all")
			test.GET("/:id")
			test.POST("/create")
		}
	}

	return router
}
