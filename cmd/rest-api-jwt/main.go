package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "testForJun02/docs"
	"testForJun02/internal/controllers"
)

//	@title TestForJun02 API
//	@version 1.0
//	@description This is a sample server.
//
// @host localhost:8080
// @BasePath /

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// @Router /adduser [post]
	e.POST("/adduser", controllers.Add, controllers.CheckAuth)

	// @Router /user [get]
	e.GET("/user", controllers.GetById, controllers.CheckAuth)

	// @Router /users [get]
	e.GET("/users", controllers.GetAll, controllers.CheckAuth)

	// @Router /deluser [delete]
	e.DELETE("/deluser", controllers.DelById, controllers.CheckAuth)

	// @Router /login [post]
	e.POST("/login", controllers.Login)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start("localhost:8080"))

}
