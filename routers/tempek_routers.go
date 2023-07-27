package routers

import (
	"smapurv1_api/controller/crud/ms"

	"github.com/gin-gonic/gin"
)

type TempekRouter struct {
	tempekRouter ms.TempekController
}

func NewRouteTempekController(tempekRouter ms.TempekController) TempekRouter {
	return TempekRouter{tempekRouter}
}

func (ar *TempekRouter) TempekRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/tempek")
	router.Use()
	router.POST("/", ar.tempekRouter.CreateTempek)
	router.GET("/", ar.tempekRouter.FindTempeks)
	router.PUT("/:tempekId", ar.tempekRouter.UpdateTempek)
	router.GET("/:tempekId", ar.tempekRouter.FindTempekById)
	router.DELETE("/:tempekId", ar.tempekRouter.DeleteTempek)
}
