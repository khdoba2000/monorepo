package auth_handler

import (
	"log"
	"net/http"
	// "monorepo/src/api_gateway/ci"
)

type AuthHandlers interface {
	TestHandler(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
}

// New creates auth handlers
func New(logger *log.Logger) AuthHandlers {
	return &authHandler{}
}

func (ah *authHandler) TestHandler(w http.ResponseWriter, r *http.Request) {
	// ah.Logger.Print("Got a request.")
	w.Write([]byte("Hello, World1!"))
}

func (ah *authHandler) SendSMS(w http.ResponseWriter, r *http.Request) {
	// container := ci.Get()
}
