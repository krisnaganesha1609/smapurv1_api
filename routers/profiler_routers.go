package routers

import (
	"github.com/gin-gonic/gin"

	"smapurv1_api/controller/auth"
	"smapurv1_api/middleware"
)

type ProfilerRouter struct {
	profilerRouter auth.ProfilerController
}

func NewRouteProfilerController(profilerRouter auth.ProfilerController) ProfilerRouter {
	return ProfilerRouter{profilerRouter}
}

func (pr *ProfilerRouter) ProfilerRoutes(rg *gin.RouterGroup) {
	router := rg.Group("profiler")
	router.Use(middleware.DeserializeUser())
	router.GET("/me", pr.profilerRouter.GetMe)
	router.PUT("/updateprofile", pr.profilerRouter.UpdateProfile)
	router.PATCH("/updatepassword", pr.profilerRouter.UpdatePassword)
}
