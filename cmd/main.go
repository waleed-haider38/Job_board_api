package main

import (
	"job-board-api/config"
	"job-board-api/controllers"
	"job-board-api/middleware"
	"job-board-api/migrations"

	"github.com/labstack/echo/v4"
)

func main() {

	// 1 Database connect
	config.ConnectDatabase()

	// 2 AutoMigrate RUN (THIS WAS MISSING)
	migrations.RunMigrations()

	// 3 Echo server
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome to Job Board API")
	})
	//Register Route
	e.POST("/register", controllers.Register)
	//Login Route
	e.POST("/login",controllers.Login)
	//Test middleare
	e.GET("/profile", controllers.Profile, middleware.JWTMiddleware)

	
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "API is running")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
