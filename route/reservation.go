package route

import (
	ctl "food-warrior/controller"

	"github.com/labstack/echo/v4"
)

func reservation(e *echo.Echo) {
	group := e.Group("/reservations")

	// List
	group.GET("", ctl.ReservationList)

	// Create
	group.POST("", ctl.ReservationCreate)

	// Delete All
	group.DELETE("", ctl.ReservationDelete)

	// Delete
	group.DELETE("/:id", ctl.ReservationDeleteByID)
}
