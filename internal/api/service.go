package api

import (
	"github.com/drizzleent/emplyees/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.ApiService
}

func NewHandler(service service.ApiService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	changeEmp := router.Group("/")
	{
		changeEmp.POST("/", h.Create)
		changeEmp.DELETE("/:id", h.Delete)
		changeEmp.PUT("/:id", h.Update)
	}

	getEmp := router.Group("/company")
	{
		getEmp.GET("/:company_id", h.GetWithCompany)

		departamentEmp := router.Group("/departament")
		{
			departamentEmp.GET("/:departament_id", h.GetWithDepartament)
		}
	}

	return router
}
