package main

import (
	"log"

	"github.com/MikeMwita/savannah-ordermanagement/internal/core/models"
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/repository"
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/usecase"
)

func main() {
	// Initialize the database
	repository.InitDB()

	// Create an example order (you can replace this with your actual order creation logic)
	order := models.Order{
		CustomerID: 123,     // Example customer ID
		Item:       "Shirt", // Example item
		Amount:     25.99,   // Example amount
	}

	// Send an SMS alert for the newly created order
	err := usecase.SendSmsAlert(order)
	if err != nil {
		log.Fatalf("failed to send SMS alert: %v", err)
	}

	log.Println("SMS alert sent successfully!")
}
