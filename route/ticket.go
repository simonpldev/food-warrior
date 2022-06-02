package route

import (
	ctl "food-warrior/controller"

	"github.com/labstack/echo/v4"
)

func reservation(e *echo.Echo) {
	group := e.Group("/tickets")

	// List
	group.GET("", ctl.TicketList)

	// Detail
	group.GET("/:id", ctl.TicketDetail)

	// Create
	group.POST("", ctl.TicketCreate)

	// Update
	group.PUT("/:id", ctl.TicketUpdate)

	// Delete All
	group.DELETE("", ctl.TicketDelete)

	// Delete
	group.DELETE("/:id", ctl.TicketDeleteByID)
}
