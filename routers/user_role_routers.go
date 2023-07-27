package routers

import (
	"smapurv1_api/controller/crud/ad"

	"github.com/gin-gonic/gin"
)

type UserRoleRouter struct {
	userRoleRouter ad.UserRoleController
}

func NewRouteUserRoleController(userRoleRouter ad.UserRoleController) UserRoleRouter {
	return UserRoleRouter{userRoleRouter}
}

func (urr *UserRoleRouter) UserRoleRoutes(rg *gin.RouterGroup) {
	router := rg.Group("user-role")
	router.Use()
	router.POST("/", urr.userRoleRouter.CreateUserRole)
	router.GET("/", urr.userRoleRouter.FindUserRoles)
	router.PUT("/:userId/:roleId", urr.userRoleRouter.UpdateUserRole)
	router.GET("/:userId/:roleId", urr.userRoleRouter.FindUserRoleById)
	router.DELETE("/:userId/:roleId", urr.userRoleRouter.DeleteUserRole)
}
