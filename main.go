package main

import (
	"github.com/MikeMwita/savannah-ordermanagement/config"
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/repository"
	"github.com/MikeMwita/savannah-ordermanagement/internal/routes"
	"github.com/MikeMwita/savannah-ordermanagement/pkg"
	"github.com/MikeMwita/savannah-ordermanagement/pkg/utils"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	_, err = config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pkg.InitDB()

	orderRepo := repository.NewOrderRepo(pkg.DB)
	customerRepo := repository.NewCustomerRepo(pkg.DB)

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register authentication and protected routes
	routes.RegisterRoutes(mux, orderRepo, customerRepo)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5556"
	}

	log.Printf("Server is running on port %s...\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
