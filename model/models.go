package model

type Food struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type FoodPayload struct {
	FoodString string `json:"food_string"`
}

type Reservation struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	ReservedFoodID []int  `json:"reserved_food_id"`
}

type ReservationPayload struct {
	Username       string `json:"username"`
	ReservedFoodID []int  `json:"reserved_food_id"`
}
