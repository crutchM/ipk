package handler

import (
	"github.com/gin-gonic/gin"
	"ipk/pkg/service"
)

//головной объект хендлера запросов, вся его суть просто навешивать методы на каждый путь и передавать им контекст gin
type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	//очевидно, группа запросов на авторизацию и регистрацию
	auth := router.Group("/auth")
	{
		auth.GET("/", h.check)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
		auth.GET("/getall", h.getAll)
	}
	//основной набор методов апи userIdentity-просто метод, который проверяет валидность jwt токена, полученного после авторизации
	api := router.Group("/api", h.userIdentity)
	{
		chair := api.Group("/chair")
		{
			//вспомогательные методы для кафедр, возможно не буду использоваться
			chair.GET("/getall", h.getAllChairs)
			chair.POST("/create", h.createChair)
		}

		test := api.Group("/test")
		{
			//методв create-задел на будущее, пока необходим только один вариант опроса
			test.GET("/:id", h.GetTest)
			test.POST("/create", h.CreateTest)
			test.POST("/sendResults", h.SendResult)
			test.POST("/sendStat", h.SendStat)
		}
		stat := api.Group("/stat")
		{
			stat.GET("/getStat", h.getStat)
			stat.GET("/getIndividual/:id", h.getStatByTeacher)
		}
	}

	return router
}
