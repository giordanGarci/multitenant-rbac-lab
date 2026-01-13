package interceptors

import (
	"context"
	"net/http"
)

type UserContext struct {
	UserID string
}

type contextKey string

const userKey contextKey = "user_data"

func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User-ID")

		if userID == "" {
			http.Error(w, "X-User-ID header missing", http.StatusUnauthorized)
			return
		}

		userCtx := &UserContext{UserID: userID}
		ctx := r.Context()
		ctx = context.WithValue(ctx, userKey, userCtx)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
func FromContext(ctx context.Context) (UserContext, bool) {
	u, ok := ctx.Value(userKey).(UserContext)
	return u, ok
}
