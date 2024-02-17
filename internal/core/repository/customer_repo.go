package repository

import (
	"database/sql"
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/adapters"
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/models"
)

type CustomerRepo struct {
	db *sql.DB
}

func (c CustomerRepo) AddCustomer(customer models.Customer) error {
	// Prepare the SQL statement to insert the customer data
	stmt, err := c.db.Prepare("INSERT INTO customers (name, code, phone, email) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement to insert the customer data
	_, err = stmt.Exec(customer.Name, customer.Code, customer.Phone, customer.Email)
	if err != nil {
		return err
	}

	return nil
}

// GetCustomerByID queries the database for a customer by its ID and returns a customer object
func (c CustomerRepo) GetCustomerByID(customerID int) (models.Customer, error) {
	// Prepare a SQL statement to select the customer data
	stmt, err := c.db.Prepare("SELECT id, name, phone FROM customers WHERE id = $1")
	if err != nil {
		return models.Customer{}, err
	}
	defer stmt.Close()

	// Execute the statement and scan the customer data
	var customer models.Customer
	err = stmt.QueryRow(customerID).Scan(&customer.ID, &customer.Name, &customer.Phone)
	if err != nil {
		return models.Customer{}, err
	}

	// Return the customer object and nil error
	return customer, nil
}

// NewCustomerRepo creates a new customer repository with a given database connection
func NewCustomerRepo(db *sql.DB) adapters.CustomerRepository {
	return &CustomerRepo{db: db}
}
