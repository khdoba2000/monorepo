package routers

import (
	"monorepo/src/api_gateway/handlers"
	"monorepo/src/api_gateway/pkg/tracing"
	"net/http"
)

func RegisterCustomerRoutes(r *tracing.TracedServeMux, h *handlers.Handlers) {
	r.Handle("/customer/test2", http.HandlerFunc(h.CustomerHandlers.TestHandler2))
}
