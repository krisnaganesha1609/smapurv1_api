package routers

import (
	"github.com/gin-gonic/gin"

	"smapurv1_api/controller/auth"
	"smapurv1_api/middleware"
)

type AuthRouter struct {
	authRouter auth.AuthController
}

func NewRouteAuthController(authRouter auth.AuthController) AuthRouter {
	return AuthRouter{authRouter}
}

func (ar *AuthRouter) AuthRoutes(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("/by-username", ar.authRouter.AuthenticateByUsername)
	router.POST("/by-nik", ar.authRouter.AuthenticateByNIK)
	router.POST("/logout", middleware.DeserializeUser(), ar.authRouter.Logout)
	router.GET("/verifyemail/:verificationCode", ar.authRouter.VerifyEmail)
	router.POST("/forgotpassword", ar.authRouter.ForgotPassword)
	router.PATCH("/resetpassword/:resetToken", ar.authRouter.ResetPassword)
}
