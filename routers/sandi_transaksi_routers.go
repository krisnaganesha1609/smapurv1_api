package routers

import (
	"smapurv1_api/controller/crud/ms"

	"github.com/gin-gonic/gin"
)

type SandiTransaksiRouter struct {
	sandiTransaksiRouter ms.SandiTransaksiController
}

func NewRouteSandiTransaksiController(sandiTransaksiRouter ms.SandiTransaksiController) SandiTransaksiRouter {
	return SandiTransaksiRouter{sandiTransaksiRouter}
}

func (ar *SandiTransaksiRouter) SandiTransaksiRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/sanditransaksi")
	router.Use()
	router.POST("/", ar.sandiTransaksiRouter.CreateSandiTransaksi)
	router.GET("/", ar.sandiTransaksiRouter.FindSandiTransaksis)
	router.PUT("/:sandiTransaksiId", ar.sandiTransaksiRouter.UpdateSandiTransaksi)
	router.GET("/:sandiTransaksiId", ar.sandiTransaksiRouter.FindSandiTransaksiById)
	router.DELETE("/:sandiTransaksiId", ar.sandiTransaksiRouter.DeleteSandiTransaksi)
}
