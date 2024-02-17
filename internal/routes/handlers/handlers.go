package handlers

import (
	"encoding/json"
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/adapters"
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/models"
	"net/http"
)

type Handler struct {
	orderRepo    adapters.OrderRepository
	customerRepo adapters.CustomerRepository
	smsRepo      adapters.SMSRepository
}

func (h *Handler) AddCustomerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Add the customer to the database
	err = h.customerRepo.AddCustomer(customer)
	if err != nil {
		http.Error(w, "Failed to add customer to the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Customer added successfully"))
}

func (h *Handler) AddOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Add the order to the database
	_, err = h.orderRepo.CreateOrder(order)
	if err != nil {
		http.Error(w, "Failed to add order to the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Order added successfully"))
}

func NewHandler(orderRepo adapters.OrderRepository, customerRepo adapters.CustomerRepository, smsRepo adapters.SMSRepository) *Handler {
	return &Handler{
		orderRepo:    orderRepo,
		customerRepo: customerRepo,
		smsRepo:      smsRepo,
	}
}
