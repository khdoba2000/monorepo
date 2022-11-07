package routers

import (
	"monorepo/src/api_gateway/handlers"
	"net/http"
)

func RegisterAuthRoutes(mux *http.ServeMux, h *handlers.Handlers) {
	mux.HandleFunc("/test1", h.AuthHandlers.TestHandler)
}
