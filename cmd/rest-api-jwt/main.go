// Package main
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"testForJun02/internal/controllers"
)

//	@title TestForJun02 API
//	@version 1.0
//	@description This is a sample server celler server.

// @host localhost:8080
// @BasePath /
func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// GET GetById method to get user by id. You must send json with parameters: id = ?
	// Example {
	//    "id": 1
	// }
	e.GET("/user", controllers.GetById)

	// GET method to get users.
	e.GET("/users", controllers.GetAll)

	// POST method to add user. You must send json with parameters: name, surname, birthday
	// Example {
	//    "name":"Ivan",
	//    "surname":"Ivanov",
	//    "birthday": "1991-11-16"
	// }
	e.POST("/adduser", controllers.Add)

	// DELETE method to delete user. You must send json with parameters: id = ?
	// Example {
	//    "id": 1
	// }
	e.DELETE("/deluser", controllers.DelById)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start("localhost:8080"))

}
