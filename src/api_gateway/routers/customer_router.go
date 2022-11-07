package routers

import (
	"monorepo/src/api_gateway/handlers"
	"net/http"
)

func RegisterCustomerRoutes(mux *http.ServeMux, h *handlers.Handlers) {
	mux.HandleFunc("/test2", h.CustomerHandlers.TestHandler2)
}
