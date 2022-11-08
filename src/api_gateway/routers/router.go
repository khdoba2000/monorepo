package routers

import (
	"monorepo/src/api_gateway/handlers"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(mux *mux.Router, h *handlers.Handlers) {
	mux.HandleFunc("/test1", h.AuthHandlers.TestHandler)
}
