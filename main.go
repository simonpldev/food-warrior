package main

import (
	"food-warrior/db"
	"food-warrior/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	db.Connect()
}

func main() {
	//Echo
	e := echo.New()

	//CORS
	e.Use(middleware.CORS())

	//Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}  -->  method=${method}, host=${host}, uri=${uri}, status=${status}\n",
	}))

	//Add routers
	route.Bootstrap(e)

	//Open port
	e.Start(":9000")
}
