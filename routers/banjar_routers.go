package routers

import (
	"smapurv1_api/controller/crud/ms"

	"github.com/gin-gonic/gin"
)

type BanjarRouter struct {
	banjarRouter ms.BanjarController
}

func NewRouteBanjarController(banjarRouter ms.BanjarController) BanjarRouter {
	return BanjarRouter{banjarRouter}
}

func (ar *BanjarRouter) BanjarRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/banjar")
	router.Use()
	router.POST("/", ar.banjarRouter.CreateBanjar)
	router.GET("/", ar.banjarRouter.FindBanjars)
	router.PUT("/:banjarId", ar.banjarRouter.UpdateBanjar)
	router.GET("/:banjarId", ar.banjarRouter.FindBanjarById)
	router.DELETE("/:banjarId", ar.banjarRouter.DeleteBanjar)
}
