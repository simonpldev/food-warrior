package model

type Food struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type FoodCreatePayload struct {
	FoodString string `json:"food_string"`
}

type FoodUpdatePayload struct {
	Name string `json:"name"`
}

type Ticket struct {
	ID           int           `json:"id"`
	Username     string        `json:"username"`
	Reservations []Reservation `json:"reservations"`
}

type TicketPayload struct {
	Username     string               `json:"username"`
	Reservations []ReservationPayload `json:"reservations"`
}

type Reservation struct {
	ID       int `json:"id"`
	TicketID int `json:"ticket_id"`
	FoodID   int `json:"food_id"`
}

type ReservationPayload struct {
	TicketID int `json:"ticket_id"`
	FoodID   int `json:"food_id"`
}
