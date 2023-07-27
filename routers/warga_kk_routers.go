package routers

import (
	"smapurv1_api/controller/crud/tr"

	"github.com/gin-gonic/gin"
)

type WargaKKRouter struct {
	wargaKKRouter tr.WargaKKController
}

func NewRouteWargaKKController(wargaKKRouter tr.WargaKKController) WargaKKRouter {
	return WargaKKRouter{wargaKKRouter}
}

func (wkr *WargaKKRouter) WargaKKRoutes(rg *gin.RouterGroup) {
	router := rg.Group("wargakk")
	router.Use()
	router.POST("/", wkr.wargaKKRouter.CreateWargaKK)
	router.GET("/", wkr.wargaKKRouter.FindWargaKKs)
	router.PUT("/:wargaId/:nik", wkr.wargaKKRouter.UpdateWargaKK)
	router.GET("/:wargaId/:nik", wkr.wargaKKRouter.FindWargaKKById)
	router.DELETE("/:wargaId/:nik", wkr.wargaKKRouter.DeleteWargaKK)
}
