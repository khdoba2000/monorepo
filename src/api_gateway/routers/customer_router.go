package routers

import (
	"monorepo/src/api_gateway/handlers"

	"github.com/gorilla/mux"
)

func RegisterCustomerRoutes(mux *mux.Router, h *handlers.Handlers) {
	mux.HandleFunc("/test2", h.CustomerHandlers.TestHandler2)
}
