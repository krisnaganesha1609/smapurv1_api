package routers

import (
	"smapurv1_api/controller/crud/ms"

	"github.com/gin-gonic/gin"
)

type AgamaRouter struct {
	agamaRouter ms.AgamaController
}

func NewRouteAgamaController(agamaRouter ms.AgamaController) AgamaRouter {
	return AgamaRouter{agamaRouter}
}

func (ar *AgamaRouter) AgamaRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/agama")
	router.Use()
	router.POST("/", ar.agamaRouter.CreateAgama)
	router.GET("/", ar.agamaRouter.FindAgamas)
	router.PUT("/:agamaId", ar.agamaRouter.UpdateAgama)
	router.GET("/:agamaId", ar.agamaRouter.FindAgamaById)
	router.DELETE("/:agamaId", ar.agamaRouter.DeleteAgama)
}
