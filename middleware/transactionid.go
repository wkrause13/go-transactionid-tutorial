package middleware

import (
	"context"
	"math/rand"
	"net/http"
	"time"
)

type key int

const (
	TransactionId key = iota
)

func TransactionIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//would use a uuid in prod, want to keep things simple here
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		ctx := context.WithValue(r.Context(), TransactionId, r1.Intn(1000000000))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
