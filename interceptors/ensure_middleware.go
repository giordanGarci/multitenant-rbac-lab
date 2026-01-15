package interceptors

import (
	"net/http"

	"github.com/giordanGarci/api-tenants/structs"
)

// Estrutura do EnsureRole
func EnsureRole(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := FromContext(r.Context())
			if !ok {
				http.Error(w, "user not found in context", http.StatusUnauthorized)
				return
			}

			hasRole := false
			for _, role := range allowedRoles {
				if user.Role == structs.Role(role) {
					hasRole = true
					break
				}
			}

			if !hasRole {
				http.Error(w, "forbidden: insufficient role", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
