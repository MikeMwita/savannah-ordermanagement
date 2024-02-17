package repository

import (
	"database/sql"
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/adapters"
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/models"
)

type OrderRepo struct {
	db *sql.DB
}

// CreateOrder inserts a new order into the database and returns its ID
func (o OrderRepo) CreateOrder(order models.Order) (int, error) {
	// Prepare a SQL statement to insert the order data
	stmt, err := o.db.Prepare("INSERT INTO orders (customer_id, item, amount, time) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute the statement and get the order ID
	var orderID int
	err = stmt.QueryRow(order.CustomerID, order.Item, order.Amount, order.Time).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	// Return the order ID and nil error
	return orderID, nil
}

// GetCustomerByID queries the database for a customer by its ID and returns a customer object
func (o OrderRepo) GetCustomerByID(customerID int) (models.Customer, error) {
	// Prepare a SQL statement to select the customer data
	stmt, err := o.db.Prepare("SELECT id, name, phone FROM customers WHERE id = $1")
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

// NewOrderRepo creates a new order repository with a given database connection
func NewOrderRepo(db *sql.DB) adapters.OrderRepository {
	return &OrderRepo{db: db}
}
