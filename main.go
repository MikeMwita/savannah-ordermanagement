package main

import (
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/repository"
	"github.com/MikeMwita/savannah-ordermanagement/internal/routes"
	"log"
	"net/http"
)

func main() {

	//_, err := config.LoadConfig()
	//if err != nil {
	//	log.Fatal("Loading config failed", err)
	//}

	// Initialize the database
	repository.InitDB()

	orderRepo := repository.NewOrderRepo(repository.DB)
	customerRepo := repository.NewCustomerRepo(repository.DB)

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register authentication and protected routes
	routes.RegisterRoutes(mux, orderRepo, customerRepo)

	// Start the server
	port := ":5556"
	log.Printf("Server is running on port %s...\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
