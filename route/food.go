package route

import (
	ctl "food-warrior/controller"

	"github.com/labstack/echo/v4"
)

func food(e *echo.Echo) {
	group := e.Group("/foods")

	// List
	group.GET("", ctl.FoodList)

	// Create
	group.POST("", ctl.FoodCreate)

	// Delete All
	group.DELETE("", ctl.FoodDelete)

	// Delete
	group.DELETE("/:id", ctl.FoodDeleteByID)
}
