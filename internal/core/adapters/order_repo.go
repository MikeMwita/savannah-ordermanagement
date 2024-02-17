package adapters

import "github.com/MikeMwita/savannah-ordermanagement/internal/core/models"

type OrderRepository interface {
	CreateOrder(order models.Order) (int, error)
	GetCustomerByID(customerID int) (models.Customer, error)
}

type CustomerRepository interface {
	GetCustomerByID(customerID int) (models.Customer, error)
	AddCustomer(customer models.Customer) error
}

type SMSRepository interface {
	SendSMS(message string, phone string) error
}
