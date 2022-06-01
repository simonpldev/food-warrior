package controller

import (
	"food-warrior/db"
	"food-warrior/model"
	"food-warrior/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ReservationList(c echo.Context) error {

	rl := db.GetReservationList()

	rs := convertToResponse(rl)

	//Return 200
	return util.Response200(c, "", rs)
}

func convertToResponse(reservationList []model.Reservation) []model.ReservationResponse {
	var res []model.ReservationResponse
	for _, r := range reservationList {
		food := db.GetFoodByID(r.ReservedFoodID)
		rr := model.ReservationResponse{
			ID:           r.ID,
			Username:     r.Username,
			ReservedFood: food,
		}
		res = append(res, rr)
	}
	return res
}

func ReservationCreate(c echo.Context) error {
	var (
		payload model.ReservationPayload
	)

	//Bind and parse to struct
	if err := c.Bind(&payload); err != nil {
		return util.Response400(c, err.Error())
	}

	db.AddReservation(payload)

	//Return 200
	return util.Response200(c, "", nil)
}

func ReservationDeleteByID(c echo.Context) error {
	var idStr = c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Delete reservation list
	db.DeleteReservationByID(id)

	//Return 200
	return util.Response200(c, "", nil)
}

func ReservationDelete(c echo.Context) error {

	//Delete reservation list
	db.DeleteAllReservation()

	//Return 200
	return util.Response200(c, "", nil)
}
