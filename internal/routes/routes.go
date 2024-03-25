package routes

import (
	"fmt"
	"github.com/MikeMwita/savannah-ordermanagement/internal/core/adapters"
	"github.com/MikeMwita/savannah-ordermanagement/internal/routes/handlers"
	"github.com/MikeMwita/savannah-ordermanagement/pkg/authenticator"
	"github.com/gorilla/sessions"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"
)

// RegisterRoutes --->registers all routes for the application
func RegisterRoutes(mux *http.ServeMux, orderRepo adapters.OrderRepository, customerRepo adapters.CustomerRepository, auth *authenticator.Authenticator) {
	handler := handlers.NewHandler(orderRepo, customerRepo)

	store := sessions.NewCookieStore([]byte("mike"))
	mux.Handle("/", otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the Savannah OrderManagement API!")
	}), "Index"))

	mux.Handle("/login", otelhttp.NewHandler(handlers.LoginHandler(auth, store), "Login"))
	mux.Handle("/logout", otelhttp.NewHandler(http.HandlerFunc(handlers.LogoutHandler), "Logout"))
	mux.Handle("/callback", otelhttp.NewHandler(http.HandlerFunc(handlers.CallBackHandler(auth, store)), "Callback"))
	mux.Handle("/user", otelhttp.NewHandler(http.HandlerFunc(handlers.UserHandler), "User"))

	//mux.HandleFunc("/user", middleware.IsAuthenticated(user.Handler))
	mux.Handle("/signup", otelhttp.NewHandler(http.HandlerFunc(handler.AddCustomerHandler), "AddCustomer"))
	mux.Handle("/order", otelhttp.NewHandler(http.HandlerFunc(handler.AddOrderHandler), "AddOrder"))

}
