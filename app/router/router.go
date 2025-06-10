package router

import (
	"github.com/labstack/echo/v4"
	"app/controller"
)

func SetupRoutes(e *echo.Echo) {
	homeController := &controller.HomeController{}
	e.GET("/", homeController.Index)

	e.GET("/hello-golang", homeController.HelloGolang)
	e.GET("/rest-api-echo", homeController.RestAPI)
	e.GET("/db-connection", homeController.DBConnection)
	e.GET("/jwt-auth", homeController.JWTAuth)
	e.GET("/mvc-pattern", homeController.MVCPattern)
	e.GET("/deploy-https", homeController.DeployHTTPS)
}
