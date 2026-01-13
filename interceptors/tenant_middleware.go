package interceptors

import (
	"context"
	"net/http"
)

type contexKey string

const tenantKey contexKey = "tenant_id"

func TenantMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenantID := r.Header.Get("X-Tenant-ID")

		if tenantID == "" {
			http.Error(w, "X-Tenant-ID header missing", http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, tenantKey, tenantID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
func TenantFromContext(ctx context.Context) (string, bool) {
	t, ok := ctx.Value(tenantKey).(string)
	return t, ok
}
