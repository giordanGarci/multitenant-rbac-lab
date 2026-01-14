package interceptors

import (
	"context"
	"net/http"

	"github.com/giordanGarci/api-tenants/structs"
)

type contextKey string

const userKey contextKey = "user_data"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User-ID")
		tenantID := r.Header.Get("X-Tenant-ID")
		role := r.Header.Get("X-User-Role")

		if userID == "" || tenantID == "" || role == "" {
			http.Error(w, "X-User-ID, X-Tenant-ID or X-User-Role header missing", http.StatusUnauthorized)
			return
		}

		userCtx := &structs.UserContext{UserID: userID, TenantID: tenantID, Role: structs.Role(role)}
		ctx := r.Context()
		ctx = context.WithValue(ctx, userKey, userCtx)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
func FromContext(ctx context.Context) (structs.UserContext, bool) {
	u, ok := ctx.Value(userKey).(*structs.UserContext)
	return *u, ok
}
