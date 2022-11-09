package customer_handler

import (
	"log"
	"monorepo/src/api_gateway/utils"
	libs "monorepo/src/libs/utils"
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

func (ch *customerHandler) SendCode(w http.ResponseWriter, r *http.Request) {
	var body utils.ReqSendCode
	if err := utils.BodyParser(r, body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.WriteJSON(w, utils.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if body.Phone != "" {
		if err := libs.SendCodeToPhone(body.Phone); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			utils.WriteJSON(w, utils.Error{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
	} else if body.Email != "" {
		if err := libs.SendCodeToEmail(body.Email); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			utils.WriteJSON(w, utils.Error{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
	} else if body.Email == "" && body.Phone == "" {
		w.WriteHeader(http.StatusBadRequest)
		utils.WriteJSON(w, utils.Error{
			Status:  http.StatusBadRequest,
			Message: "Data is not provided",
		})
		return
	}

}

func (ch *customerHandler) VerfyCode(w http.ResponseWriter, r *http.Request) {
	var body utils.ReqCheckCode
	if err := utils.BodyParser(r, body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.WriteJSON(w, utils.Error{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

}
