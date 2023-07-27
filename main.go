package main

import (
	"log"
	"net/http"
	"smapurv1_api/controller/auth"
	ad "smapurv1_api/controller/crud/ad"
	ms "smapurv1_api/controller/crud/ms"
	tr "smapurv1_api/controller/crud/tr"
	r "smapurv1_api/routers"
	s "smapurv1_api/setup"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	//Create New Controller and Router Variables

	AuthController auth.AuthController
	AuthRouter     r.AuthRouter

	MenuController ad.MenuController
	MenuRouter     r.MenuRouter

	RoleController ad.RoleController
	RoleRouter     r.RoleRouter

	UserController ad.UserController
	UserRouter     r.UserRouter

	UserRoleController ad.UserRoleController
	UserRoleRouter     r.UserRoleRouter

	RoleMenuController ad.RoleMenuController
	RoleMenuRouter     r.RoleMenuRouter

	AgamaController ms.AgamaController
	AgamaRouter     r.AgamaRouter

	BanjarController ms.BanjarController
	BanjarRouter     r.BanjarRouter

	HubKeluargaController ms.HubKeluargaController
	HubKeluargaRouter     r.HubKeluargaRouter

	PendidikanController ms.PendidikanController
	PendidikanRouter     r.PendidikanRouter

	TempekController ms.TempekController
	TempekRouter     r.TempekRouter

	SandiTransaksiController ms.SandiTransaksiController
	SandiTransaksiRouter     r.SandiTransaksiRouter

	WargaController tr.WargaController
	WargaRouter     r.WargaRouter

	WargaKKController tr.WargaKKController
	WargaKKRouter     r.WargaKKRouter

	WargaTransaksiController tr.WargaTransaksiController
	WargaTransaksiRouter     r.WargaTransaksiRouter
)

func init() {
	config, err := s.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	//Init Database Connectivity
	s.ConnectDatabase(&config)

	//Insert New Controller and Router Initializers

	AuthController = auth.NewAuthController(s.DB)
	AuthRouter = r.NewRouteAuthController(AuthController)

	MenuController = ad.NewMenuController(s.DB)
	MenuRouter = r.NewRouteMenuController(MenuController)

	RoleController = ad.NewRoleController(s.DB)
	RoleRouter = r.NewRouteRoleController(RoleController)

	UserController = ad.NewUserController(s.DB)
	UserRouter = r.NewRouteUserController(UserController)

	UserRoleController = ad.NewUserRoleController(s.DB)
	UserRoleRouter = r.NewRouteUserRoleController(UserRoleController)

	RoleMenuController = ad.NewRoleMenuController(s.DB)
	RoleMenuRouter = r.NewRouteRoleMenuController(RoleMenuController)

	AgamaController = ms.NewAgamaController(s.DB)
	AgamaRouter = r.NewRouteAgamaController(AgamaController)

	BanjarController = ms.NewBanjarController(s.DB)
	BanjarRouter = r.NewRouteBanjarController(BanjarController)

	HubKeluargaController = ms.NewHubKeluargaController(s.DB)
	HubKeluargaRouter = r.NewRouteHubKeluargaController(HubKeluargaController)

	PendidikanController = ms.NewPendidikanController(s.DB)
	PendidikanRouter = r.NewRoutePendidikanController(PendidikanController)

	TempekController = ms.NewTempekController(s.DB)
	TempekRouter = r.NewRouteTempekController(TempekController)

	SandiTransaksiController = ms.NewSandiTransaksiController(s.DB)
	SandiTransaksiRouter = r.NewRouteSandiTransaksiController(SandiTransaksiController)

	WargaController = tr.NewWargaController(s.DB)
	WargaRouter = r.NewRouteWargaController(WargaController)

	WargaKKController = tr.NewWargaKKController(s.DB)
	WargaKKRouter = r.NewRouteWargaKKController(WargaKKController)

	WargaTransaksiController = tr.NewWargaTransaksiController(s.DB)
	WargaTransaksiRouter = r.NewRouteWargaTransaksiController(WargaTransaksiController)

	server = gin.Default()

}

func main() {
	config, err := s.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	//Migrate The Tables
	s.Migrations()

	//Enable the Default CORS Configuration
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(c *gin.Context) {
		message := "SMAPUR RESTAPI v1.0"
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	//Connect the routes
	AuthRouter.AuthRoutes(router)
	MenuRouter.MenuRoutes(router)
	RoleRouter.RoleRoutes(router)
	UserRouter.UserRoutes(router)
	UserRoleRouter.UserRoleRoutes(router)
	RoleMenuRouter.RoleMenuRoutes(router)
	AgamaRouter.AgamaRoutes(router)
	BanjarRouter.BanjarRoutes(router)
	HubKeluargaRouter.HubKeluargaRoutes(router)
	PendidikanRouter.PendidikanRoutes(router)
	TempekRouter.TempekRoutes(router)
	SandiTransaksiRouter.SandiTransaksiRoutes(router)
	WargaRouter.WargaRoutes(router)
	WargaKKRouter.WargaKKRoutes(router)
	WargaTransaksiRouter.WargaTransaksiRoutes(router)

	//Run Router At Port :8000
	log.Fatal(server.Run(":" + config.ServerPort))
}
