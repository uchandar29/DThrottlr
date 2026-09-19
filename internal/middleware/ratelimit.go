package middleware

import (
	"fmt"
	"net"
	"net/http"

	"github.com/uchandar29/DThrottlr/internal/limiter"
)

func RateLimit(l *limiter.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientID := clientIDFromRequest(r)

			allow, remaining, err := l.Allow(r.Context(), clientID)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("X-RateLimit-Remaining", itoa(remaining))

			if !allow {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func clientIDFromRequest(r *http.Request) string {
	if id := r.Header.Get("X-Client-ID"); id != "" {
		return id
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

func itoa(n int) string {
	return fmt.Sprintf("%d", n)
}
