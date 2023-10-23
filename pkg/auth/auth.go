package auth

import (
	"context"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
)

var contextKey struct{}

func Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		headerString := r.Header.Get(authorizationHeader)
		r = r.WithContext(context.WithValue(r.Context(), contextKey, headerString))
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func FromContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(contextKey).(string)
	return token, ok
}
