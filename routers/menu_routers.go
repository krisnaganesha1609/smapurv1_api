package routers

import (
	"smapurv1_api/controller/crud/ad"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct {
	menuRouter ad.MenuController
}

func NewRouteMenuController(menuRouter ad.MenuController) MenuRouter {
	return MenuRouter{menuRouter}
}

func (mc *MenuRouter) MenuRoutes(rg *gin.RouterGroup) {
	router := rg.Group("menu")
	router.Use()
	router.POST("/", mc.menuRouter.CreateMenu)
	router.GET("/", mc.menuRouter.FindMenus)
	router.PUT("/:menuId", mc.menuRouter.UpdateMenu)
	router.GET("/:menuId", mc.menuRouter.FindMenuById)
	router.DELETE("/:menuId", mc.menuRouter.DeleteMenu)
}
