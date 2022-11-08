package routers

import (
	"monorepo/src/api_gateway/handlers"

	"github.com/gorilla/mux"
)

func RegisterCustomerRoutes(r *mux.Router, h *handlers.Handlers) {
	r = r.PathPrefix("/customer").Subrouter()
	r.HandleFunc("/test2", h.CustomerHandlers.TestHandler2)
}
