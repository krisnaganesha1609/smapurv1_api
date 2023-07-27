package routers

import (
	"smapurv1_api/controller/crud/ad"

	"github.com/gin-gonic/gin"
)

type RoleMenuRouter struct {
	roleMenuRouter ad.RoleMenuController
}

func NewRouteRoleMenuController(roleMenuRouter ad.RoleMenuController) RoleMenuRouter {
	return RoleMenuRouter{roleMenuRouter}
}

func (rmr *RoleMenuRouter) RoleMenuRoutes(rg *gin.RouterGroup) {
	router := rg.Group("role-menu")
	router.Use()
	router.POST("/", rmr.roleMenuRouter.CreateRoleMenu)
	router.GET("/", rmr.roleMenuRouter.FindRoleMenus)
	router.PUT("/:roleId/:menuId", rmr.roleMenuRouter.UpdateRoleMenu)
	router.GET("/:roleId/:menuId", rmr.roleMenuRouter.FindRoleMenuById)
	router.DELETE("/:roleId/:menuId", rmr.roleMenuRouter.DeleteRoleMenu)
}
