package routers

import (
	"smapurv1_api/controller/crud/ad"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct {
	roleRouter ad.RoleController
}

func NewRouteRoleController(roleRouter ad.RoleController) RoleRouter {
	return RoleRouter{roleRouter}
}

func (rc *RoleRouter) RoleRoutes(rg *gin.RouterGroup) {
	router := rg.Group("role")
	router.Use()
	router.POST("/", rc.roleRouter.CreateRole)
	router.GET("/", rc.roleRouter.FindRoles)
	router.PUT("/:roleId", rc.roleRouter.UpdateRole)
	router.GET("/:roleId", rc.roleRouter.FindRoleById)
	router.DELETE("/:roleId", rc.roleRouter.DeleteRole)
}
