package routers

import (
	"smapurv1_api/controller/crud/tr"

	"github.com/gin-gonic/gin"
)

type WargaRouter struct {
	wargaRouter tr.WargaController
}

func NewRouteWargaController(wargaRouter tr.WargaController) WargaRouter {
	return WargaRouter{wargaRouter}
}

func (wr *WargaRouter) WargaRoutes(rg *gin.RouterGroup) {
	router := rg.Group("warga")
	router.Use()
	router.POST("/", wr.wargaRouter.CreateWarga)
	router.GET("/", wr.wargaRouter.FindWargas)
	router.PUT("/:wargaId", wr.wargaRouter.UpdateWarga)
	router.GET("/:wargaId", wr.wargaRouter.FindWargaById)
	router.DELETE("/:wargaId", wr.wargaRouter.DeleteWarga)
}
