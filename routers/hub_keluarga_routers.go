package routers

import (
	"smapurv1_api/controller/crud/ms"

	"github.com/gin-gonic/gin"
)

type HubKeluargaRouter struct {
	hubKeluargaRouter ms.HubKeluargaController
}

func NewRouteHubKeluargaController(hubKeluargaRouter ms.HubKeluargaController) HubKeluargaRouter {
	return HubKeluargaRouter{hubKeluargaRouter}
}

func (ar *HubKeluargaRouter) HubKeluargaRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/hubkeluarga")
	router.Use()
	router.POST("/", ar.hubKeluargaRouter.CreateHubKeluarga)
	router.GET("/", ar.hubKeluargaRouter.FindHubKeluargas)
	router.PUT("/:hubKeluargaId", ar.hubKeluargaRouter.UpdateHubKeluarga)
	router.GET("/:hubKeluargaId", ar.hubKeluargaRouter.FindHubKeluargaById)
	router.DELETE("/:hubKeluargaId", ar.hubKeluargaRouter.DeleteHubKeluarga)
}
