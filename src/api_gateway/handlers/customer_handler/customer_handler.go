package customer_handler

import (
	"monorepo/src/libs/log"
	"net/http"
)

type CustomerHandlers interface {
	TestHandler2(w http.ResponseWriter, r *http.Request)
}

type customerHandler struct {
	logger log.Factory
}

// New creates handler
func New(logger log.Factory) CustomerHandlers {
	return &customerHandler{
		logger: logger,
	}
}
func (ch *customerHandler) TestHandler2(w http.ResponseWriter, r *http.Request) {
	// ch.logger.Print("Got a request.")
	w.Write([]byte("Hello, World2!"))
}
