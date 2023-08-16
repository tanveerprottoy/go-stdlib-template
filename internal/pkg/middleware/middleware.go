package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/response"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/jwtpkg"
)

// JSONContentTypeMiddleWare content type json setter middleware
func JSONContentTypeMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// CORSEnableMiddleWare enable cors
func CORSEnableMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		next.ServeHTTP(w, r)
	})
}

// TimeoutHandler is a middleware to add http.TimeoutHandler.
func TimeoutHandler(timeout time.Duration) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.TimeoutHandler(next, timeout, "Timeout.")
	}
}

// JWTMiddleWare checks auth of the request
func JWTMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			// Token is missing
			response.RespondError(http.StatusForbidden, "error", errors.New("auth token is missing"), w)
			return
		}
		split := strings.Split(tokenHeader, " ")
		// token format is `Bearer {tokenBody}`
		if len(split) != 2 {
			response.RespondError(http.StatusForbidden, "error", errors.New("token format is invalid"), w)
			return
		}
		tokenBody := split[1]
		claims, err := jwtpkg.VerifyToken(tokenBody)
		if err != nil {
			response.RespondError(http.StatusForbidden, "error", err, w)
			return
		}
		ctx := context.WithValue(r.Context(), constant.ContextPayloadKey, claims.Payload)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
