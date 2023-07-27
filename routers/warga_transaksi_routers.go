package routers

import (
	"smapurv1_api/controller/crud/tr"

	"github.com/gin-gonic/gin"
)

type WargaTransaksiRouter struct {
	wargaTransaksiRouter tr.WargaTransaksiController
}

func NewRouteWargaTransaksiController(wargaTransaksiRouter tr.WargaTransaksiController) WargaTransaksiRouter {
	return WargaTransaksiRouter{wargaTransaksiRouter}
}

func (wr *WargaTransaksiRouter) WargaTransaksiRoutes(rg *gin.RouterGroup) {
	router := rg.Group("warga-transaksi")
	router.Use()
	router.POST("/", wr.wargaTransaksiRouter.CreateWargaTransaksi)
	router.GET("/", wr.wargaTransaksiRouter.FindWargaTransaksis)
	router.PUT("/:wargaTransaksiId/:sandiId/:tglTrans/:serialId", wr.wargaTransaksiRouter.UpdateWargaTransaksi)
	router.GET("/:wargaTransaksiId/:sandiId/:tglTrans/:serialId", wr.wargaTransaksiRouter.FindWargaTransaksiById)
	router.DELETE("/:wargaTransaksiId/:sandiId/:tglTrans/:serialId", wr.wargaTransaksiRouter.DeleteWargaTransaksi)
}
