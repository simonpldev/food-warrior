package controller

import (
	"food-warrior/db"
	"food-warrior/model"
	"food-warrior/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

func TicketList(c echo.Context) error {

	var rs []model.Ticket

	tickets := db.GetTicketList()

	for _, t := range tickets {
		t.Reservations = db.GetReservationListByTicketID(t.ID)
		rs = append(rs, t)
	}

	//Return 200
	return util.Response200(c, "", rs)
}

func TicketDetail(c echo.Context) error {
	var idStr = c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return util.Response404(c, err.Error())
	}

	rs := db.GetTicketByID(id)
	rs.Reservations = db.GetReservationListByTicketID(id)

	//Return 200
	return util.Response200(c, "", rs)
}

func TicketUpdate(c echo.Context) error {
	var (
		idStr   = c.Param("id")
		payload model.TicketPayload
	)

	//Bind and parse to struct
	if err := c.Bind(&payload); err != nil {
		return util.Response400(c, err.Error())
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return util.Response404(c, err.Error())
	}

	db.UpdateTicketByID(id, payload)

	db.DeleteReservationByTicketID(id)
	for _, r := range payload.Reservations {
		db.AddReservation(r)
	}

	//Return 200
	return util.Response200(c, "", nil)
}

func TicketCreate(c echo.Context) error {
	var (
		payload model.TicketPayload
	)

	//Bind and parse to struct
	if err := c.Bind(&payload); err != nil {
		return util.Response400(c, err.Error())
	}

	db.AddTicket(payload)

	for _, r := range payload.Reservations {
		db.AddReservation(r)
	}

	//Return 200
	return util.Response200(c, "", nil)
}

func TicketDeleteByID(c echo.Context) error {
	var idStr = c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Delete Ticket list
	db.DeleteReservationByTicketID(id)
	db.DeleteTicketByID(id)

	//Return 200
	return util.Response200(c, "", nil)
}

func TicketDelete(c echo.Context) error {

	//Delete Ticket list
	db.DeleteAllReservation()
	db.DeleteAllTicket()

	//Return 200
	return util.Response200(c, "", nil)
}
