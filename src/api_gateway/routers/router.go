package routers

import (
	"monorepo/src/api_gateway/handlers"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router, h *handlers.Handlers) {
	r = r.PathPrefix("/auth").Subrouter()
	r.HandleFunc("/test1", h.AuthHandlers.TestHandler)
	r.HandleFunc("/sendCode/", h.AuthHandlers.SendCode)
	r.HandleFunc("/verify/", h.AuthHandlers.Verify)
	r.HandleFunc("/login/", h.AuthHandlers.Login)
}
