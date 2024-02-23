package routes

import (
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/adapters"
	"github.com/MikeMwita/savannah-ordermanagement/internal/routes/handlers"
	"net/http"
)

// RegisterRoutes registers all routes for the application
func RegisterRoutes(mux *http.ServeMux, orderRepo adapters.OrderRepository, customerRepo adapters.CustomerRepository) {
	handler := handlers.NewHandler(orderRepo, customerRepo)

	// Authentication routes
	mux.HandleFunc("/auth/google/login", handlers.LoginHandler)
	mux.HandleFunc("/auth/google/logout", handlers.LogoutHandler)
	mux.HandleFunc("/auth/google/callback", handlers.CallbackHandler)
	mux.HandleFunc("/user", handlers.UserInfoHandler)

	mux.HandleFunc("/signup", handler.AddCustomerHandler)
	mux.HandleFunc("/order", handler.AddOrderHandler)

	// Protected routes
	//mux.Handle("/signup", middleware.AuthMiddleware(http.HandlerFunc(handler.AddCustomerHandler)))
	//mux.Handle("/order", middleware.AuthMiddleware(http.HandlerFunc(handler.AddOrderHandler)))

}
