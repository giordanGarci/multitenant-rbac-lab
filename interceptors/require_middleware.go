package interceptors

import (
	"net/http"
)

// Estrutura do EnsureRole
func RequirePermission(p Permission) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := FromContext(r.Context())
			if !ok {
				http.Error(w, "user not found in context", http.StatusUnauthorized)
				return
			}

			if !HasPermission(string(user.Role), p) {
				http.Error(w, "forbidden: insufficient permission", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
