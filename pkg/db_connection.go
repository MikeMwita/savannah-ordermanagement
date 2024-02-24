package pkg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(host, port, user, password, dbname, sslmode string) {
	var err error

	// Use the provided parameters to form the connection string
	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)

	// Open the database connection
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

//
//func InitDB() error {
//	host := os.Getenv("DB_HOST")
//	port := os.Getenv("DB_PORT")
//	user := os.Getenv("DB_USER")
//	dbname := os.Getenv("DB_NAME")
//	password := os.Getenv("DB_PASSWORD")
//
//	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
//	var err error
//	DB, err = sql.Open("postgres", dbConnectionString)
//	if err != nil {
//		return fmt.Errorf("failed to connect to the database: %v", err)
//	}
//
//	// Set the connection pool limits
//	DB.SetMaxOpenConns(10)
//	DB.SetMaxIdleConns(4)
//
//	// Ensure tables are created
//	if err := CreateTables(); err != nil {
//		return err
//	}
//
//	return nil
//}

func CreateTables() error {
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
		return fmt.Errorf("could not create customers table: %v", err)
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
		return fmt.Errorf("could not create orders table: %v", err)
	}

	return nil
}
