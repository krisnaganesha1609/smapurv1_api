package routers

import (
	"smapurv1_api/controller/crud/ad"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	userRouter ad.UserController
}

func NewRouteUserController(userRouter ad.UserController) UserRouter {
	return UserRouter{userRouter}
}

func (ur *UserRouter) UserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("user")
	router.Use()
	router.POST("/", ur.userRouter.CreateUser)
	router.GET("/", ur.userRouter.FindUsers)
	router.PUT("/:userId", ur.userRouter.UpdateUser)
	router.GET("/:userId", ur.userRouter.FindUserById)
	router.DELETE("/:userId", ur.userRouter.DeleteUser)
}
