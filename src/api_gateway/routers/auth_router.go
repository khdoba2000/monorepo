package routers

import (
	"monorepo/src/api_gateway/handlers"
	"monorepo/src/api_gateway/pkg/tracing"
	"net/http"
)

func RegisterAuthRoutes(r *tracing.TracedServeMux, h *handlers.Handlers) {
	r.Handle("/auth/test1", http.HandlerFunc(h.AuthHandlers.TestHandler))
	r.Handle("/auth/staff/login", http.HandlerFunc(h.AuthHandlers.StuffLogin))
	r.Handle("/auth/staff/reset-password", http.HandlerFunc(h.AuthHandlers.ResetPassword))
	r.Handle("/auth/send-code", http.HandlerFunc(h.AuthHandlers.SendCodeHandler))
	r.Handle("/auth/verfy-code", http.HandlerFunc(h.AuthHandlers.VerfyCodeHandler))
}
