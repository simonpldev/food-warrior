package db

import (
	"database/sql"
	"food-warrior/model"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
	DB, _ = sql.Open("sqlite3", "food-warrior.db")

	stmt, _ := DB.Prepare(`
		CREATE TABLE IF NOT EXISTS "FOODS" (
			"ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"NAME" TEXT NOT NULL
		);
	`)

	stmt.Exec()

	stmt, _ = DB.Prepare(`
		CREATE TABLE IF NOT EXISTS "TICKETS" (
			"ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"USERNAME" TEXT NOT NULL
		);
	`)

	stmt.Exec()

	stmt, _ = DB.Prepare(`
		CREATE TABLE IF NOT EXISTS "RESERVATIONS" (
			"ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"TICKET_ID" INTEGER NOT NULL,
			"FOOD_ID" INTEGER NOT NULL,
			FOREIGN KEY(TICKET_ID) REFERENCES TICKETS(ID),
			FOREIGN KEY(FOOD_ID) REFERENCES FOODS(ID)
		);
	`)

	stmt.Exec()

}

func GetDB() *sql.DB {
	return DB
}

func AddFood(foodName string) {
	stmt, _ := DB.Prepare(`
		INSERT INTO FOODS(NAME) values (?)
	`)

	defer stmt.Close()

	stmt.Exec(foodName)
}

func AddTicket(t model.TicketPayload) {
	stmt, _ := DB.Prepare(`
		INSERT INTO TICKETS(USERNAME) values (?)
	`)

	defer stmt.Close()

	stmt.Exec(t.Username)
}

func AddReservation(r model.ReservationPayload) {
	stmt, _ := DB.Prepare(`
		INSERT INTO RESERVATIONS(TICKET_ID, FOOD_ID) values (?, ?)
	`)

	defer stmt.Close()

	stmt.Exec(r.TicketID, r.FoodID)
}

func DeleteTicketByID(id int) {
	stmt, _ := DB.Prepare(`
		DELETE FROM TICKETS WHERE ID = ?
	`)

	defer stmt.Close()

	stmt.Exec(id)
}

func DeleteReservationByFoodID(id int) {
	stmt, _ := DB.Prepare(`
		DELETE FROM RESERVATIONS WHERE FOOD_ID = ?
	`)

	defer stmt.Close()

	stmt.Exec(id)
}

func DeleteReservationByTicketID(id int) {
	stmt, _ := DB.Prepare(`
		DELETE FROM RESERVATIONS WHERE TICKET_ID = ?
	`)

	defer stmt.Close()

	stmt.Exec(id)
}

func DeleteFoodByID(id int) {
	stmt, _ := DB.Prepare(`
		DELETE FROM FOODS WHERE ID = ?
	`)

	defer stmt.Close()

	stmt.Exec(id)
}

func DeleteAllFood() {
	stmt, _ := DB.Prepare(`
		DELETE FROM FOODS
	`)

	defer stmt.Close()

	stmt.Exec()
}

func DeleteAllReservation() {
	stmt, _ := DB.Prepare(`
		DELETE FROM RESERVATIONS
	`)

	defer stmt.Close()

	stmt.Exec()
}

func DeleteAllTicket() {
	stmt, _ := DB.Prepare(`
		DELETE FROM TICKETS
	`)

	defer stmt.Close()

	stmt.Exec()
}

func GetFoodList() []model.Food {
	foods := []model.Food{}

	rows, _ := DB.Query(`
		SELECT * FROM FOODS
	`)

	for rows.Next() {
		var food model.Food
		rows.Scan(&food.ID, &food.Name)
		foods = append(foods, food)
	}

	return foods
}

func GetTicketList() []model.Ticket {
	tickets := []model.Ticket{}

	rows, _ := DB.Query(`
		SELECT * FROM TICKETS
	`)

	for rows.Next() {
		var ticket model.Ticket
		rows.Scan(&ticket.ID, &ticket.Username)
		tickets = append(tickets, ticket)
	}

	return tickets
}

func GetReservationListByTicketID(id int) []model.Reservation {
	reservations := []model.Reservation{}

	stmt, _ := DB.Prepare(`
		SELECT * FROM RESERVATIONS WHERE TICKET_ID = ?
	`)

	rows, _ := stmt.Query(id)

	for rows.Next() {
		var reservation model.Reservation
		rows.Scan(&reservation.ID, &reservation.TicketID, &reservation.FoodID)
		reservations = append(reservations, reservation)
	}

	return reservations
}

func GetFoodByID(id int) model.Food {
	var food model.Food

	stmt, _ := DB.Prepare(`
	SELECT * FROM FOODS WHERE ID = ?
	`)

	defer stmt.Close()

	row := stmt.QueryRow(id)

	row.Scan(&food.ID, &food.Name)

	return food
}

func GetTicketByID(id int) model.Ticket {
	var ticket model.Ticket

	stmt, _ := DB.Prepare(`
	SELECT * FROM TICKETS WHERE ID = ?
	`)

	defer stmt.Close()

	row := stmt.QueryRow(id)

	row.Scan(&ticket.ID, &ticket.Username)

	return ticket
}

func UpdateFoodByID(id int, food model.FoodUpdatePayload) {

	stmt, _ := DB.Prepare(`
	UPDATE FOODS SET NAME = ? WHERE ID = ?
	`)

	defer stmt.Close()

	stmt.Exec(food.Name, id)

}

func UpdateTicketByID(id int, ticket model.TicketPayload) {

	stmt, _ := DB.Prepare(`
	UPDATE TICKETS SET USERNAME = ? WHERE ID = ?
	`)

	defer stmt.Close()

	stmt.Exec(ticket.Username, id)

}
