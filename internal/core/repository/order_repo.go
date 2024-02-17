package db

import (
	"database/sql"
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/adapters"
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/models"
)

type OrderRepo struct {
	db *sql.DB
}

func (o OrderRepo) CreateOrder(order models.Order) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepo) GetCustomerByID(customerID int) (models.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func NewOrderRepo(db *sql.DB) adapters.OrderRepository {
	return &OrderRepo{db: db}
}
