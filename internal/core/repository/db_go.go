package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	password := "1A5hhh3qsjQQdUA6IajljFTXoDQKEcwo" // hardcoded for testing purpose only

	dbConnectionString := fmt.Sprintf("host=dpg-clk5aoeg1b2c739gus30-a.oregon-postgres.render.com port=5432 user=mike dbname=mike_dy9f password=%s sslmode=require", password)

	DB, err = sql.Open("postgres", dbConnectionString)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to the database: %v", err))
	}

	// Set the connection pool limits
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(4)

	// Ensure tables are created
	CreateTables()
}

func CreateTables() {
	createCustomersTable := `
	CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		code VARCHAR(10) UNIQUE NOT NULL,
		phone VARCHAR(15) NOT NULL
	)
	`

	_, err := DB.Exec(createCustomersTable)
	if err != nil {
		panic("Could not create customers table.")
	}

	createOrdersTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		customer_id INTEGER REFERENCES customers(id) ON DELETE CASCADE,
		item VARCHAR(50) NOT NULL,
		amount DECIMAL(10, 2) NOT NULL,
		time TIMESTAMP NOT NULL
	)
	`
	_, err = DB.Exec(createOrdersTable)
	if err != nil {
		panic("Could not create orders table.")
	}
}
