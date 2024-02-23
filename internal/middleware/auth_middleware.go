package middleware

import (
	"github.com/MikeMwita/savannah-ordermanagement/internal/routes/handlers"
	"net/http"
)

var authHandler *handlers.AuthHandler

func init() {
	authHandler = handlers.NewAuthHandler()
}

// AuthMiddleware is a middleware that checks if the user is authenticated before allowing them to access the handler function.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the ID token from the cookie
		rawIDToken, err := handlers.GetCookie(r, "id_token")
		if err != nil {
			// Redirection to the login page
			http.Redirect(w, r, "/auth/google/login", http.StatusFound)
			return
		}
		_, err = authHandler.Verifier.Verify(r.Context(), rawIDToken)
		if err != nil {
			http.Redirect(w, r, "/auth/google/login", http.StatusFound)
			return
		}
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
