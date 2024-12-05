package middleware

import (
	"net/http"
	"runtime/debug"
	"tx-parser/utils"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic with stack trace
				utils.ErrorLogger.Printf("Panic recovered: %v\n%s", err, debug.Stack())

				// Return an internal server error response
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
