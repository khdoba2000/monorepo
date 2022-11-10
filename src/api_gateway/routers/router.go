package routers

import (
	"monorepo/src/api_gateway/handlers"
	"monorepo/src/api_gateway/pkg/tracing"
	"net/http"
)

func RegisterAuthRoutes(r *tracing.TracedServeMux, h *handlers.Handlers) {
	r.Handle("/auth/test1", http.HandlerFunc(h.AuthHandlers.TestHandler))
	r.Handle("/auth/staff-login/", http.HandlerFunc(h.AuthHandlers.StuffLogin))
}
