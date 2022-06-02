package route

import (
	ctl "food-warrior/controller"

	"github.com/labstack/echo/v4"
)

func food(e *echo.Echo) {
	group := e.Group("/foods")

	// List
	group.GET("", ctl.FoodList)

	// Detail
	group.GET("/:id", ctl.FoodDetail)

	// Create
	group.POST("", ctl.FoodCreate)

	// Update
	group.PUT("/:id", ctl.FoodUpdate)

	// Delete All
	group.DELETE("", ctl.FoodDelete)

	// Delete
	group.DELETE("/:id", ctl.FoodDeleteByID)
}
