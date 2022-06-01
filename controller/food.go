package controller

import (
	"food-warrior/db"
	"food-warrior/model"
	"food-warrior/util"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func FoodList(c echo.Context) error {

	rs := db.GetFoodList()

	//Return 200
	return util.Response200(c, "", rs)
}

func FoodCreate(c echo.Context) error {
	var (
		payload model.FoodPayload
	)

	//Bind and parse to struct
	if err := c.Bind(&payload); err != nil {
		return util.Response400(c, err.Error())
	}

	foodNames := strings.Split(payload.FoodString, ",")

	if len(foodNames) == 0 {
		return util.Response400(c, "No food in payload")
	}

	for _, foodName := range foodNames {
		db.AddFoodByName(strings.TrimSpace(foodName))
	}

	//Return 200
	return util.Response200(c, "", nil)
}

func FoodDeleteByID(c echo.Context) error {
	var idStr = c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return util.Response404(c, err.Error())
	}

	//Delete food and related reservation
	db.DeleteReservationByFoodID(id)
	db.DeleteFoodByID(id)

	//Return 200
	return util.Response200(c, "", nil)
}

func FoodDelete(c echo.Context) error {

	//Delete food and reservation list
	db.DeleteAllReservation()
	db.DeleteAllFood()

	//Return 200
	return util.Response200(c, "", nil)
}
