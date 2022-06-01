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
		CREATE TABLE IF NOT EXISTS "RESERVATIONS" (
			"ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"USERNAME" TEXT NOT NULL,
			"RESERVED_FOOD_ID" INTEGER NOT NULL,
			FOREIGN KEY(RESERVED_FOOD_ID) REFERENCES FOODS(ID)
		);
	`)

	stmt.Exec()

}

func GetDB() *sql.DB {
	return DB
}

func AddFoodByName(name string) {
	stmt, _ := DB.Prepare(`
		INSERT INTO FOODS(NAME) values (?)
	`)

	defer stmt.Close()

	stmt.Exec(name)
}

func AddReservation(r model.ReservationPayload) {
	stmt, _ := DB.Prepare(`
		INSERT INTO RESERVATIONS(USERNAME, RESERVED_FOOD_ID) values (?, ?)
	`)

	defer stmt.Close()

	stmt.Exec(r.Username, r.ReservedFoodID)
}

func DeleteReservationByID(id int) {
	stmt, _ := DB.Prepare(`
		DELETE FROM RESERVATIONS WHERE ID = ?
	`)

	defer stmt.Close()

	stmt.Exec(id)
}

func DeleteReservationByFoodID(id int) {
	stmt, _ := DB.Prepare(`
		DELETE FROM RESERVATIONS WHERE RESERVED_FOOD_ID = ?
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

func GetReservationList() []model.Reservation {
	reservations := []model.Reservation{}

	rows, _ := DB.Query(`
		SELECT * FROM FOODS
	`)

	for rows.Next() {
		var reservation model.Reservation
		rows.Scan(&reservation.ID, &reservation.Username, &reservation.ReservedFoodID)
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
