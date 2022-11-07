package customer_handler

import (
	"log"
	"net/http"
)

type CustomerHandlers interface {
	TestHandler2(w http.ResponseWriter, r *http.Request)
}

type customerHandler struct {
}

// New creates handler
func New(logger *log.Logger) CustomerHandlers {
	return &customerHandler{}
}
func (ch *customerHandler) TestHandler2(w http.ResponseWriter, r *http.Request) {
	// ch.logger.Print("Got a request.")
	w.Write([]byte("Hello, World2!"))
}
