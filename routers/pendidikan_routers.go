package routers

import (
	"smapurv1_api/controller/crud/ms"

	"github.com/gin-gonic/gin"
)

type PendidikanRouter struct {
	pendidikanRouter ms.PendidikanController
}

func NewRoutePendidikanController(pendidikanRouter ms.PendidikanController) PendidikanRouter {
	return PendidikanRouter{pendidikanRouter}
}

func (ar *PendidikanRouter) PendidikanRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/pendidikan")
	router.Use()
	router.POST("/", ar.pendidikanRouter.CreatePendidikan)
	router.GET("/", ar.pendidikanRouter.FindPendidikans)
	router.PUT("/:pendidikanId", ar.pendidikanRouter.UpdatePendidikan)
	router.GET("/:pendidikanId", ar.pendidikanRouter.FindPendidikanById)
	router.DELETE("/:pendidikanId", ar.pendidikanRouter.DeletePendidikan)
}
