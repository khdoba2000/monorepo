package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleGrpcErrWithMessage(w http.ResponseWriter, err error, args ...interface{}) error {
	if err == nil {
		return nil
	}
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		// logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		WriteJSON(w, Response{Error: true,
			Data: Error{
				Status:  http.StatusInternalServerError,
				Message: st.Message(),
			}})
		return err

	} else if st.Code() == codes.NotFound {
		// logger.Error(err)
		w.WriteHeader(http.StatusNotFound)
		WriteJSON(w, Response{Error: true,
			Data: Error{
				Status:  http.StatusNotFound,
				Message: st.Message(),
			}})
		return err

	} else if st.Code() == codes.Unavailable {
		// logger.Error(err)
		w.WriteHeader(http.StatusBadGateway)
		WriteJSON(w, Response{Error: true,
			Data: Error{
				Status:  http.StatusBadGateway,
				Message: st.Message(),
			}})
		return err

	} else if st.Code() == codes.AlreadyExists {
		// logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		WriteJSON(w, Response{Error: true,
			Data: Error{
				Status:  http.StatusBadRequest,
				Message: st.Message(),
			}})
		return err

	} else if st.Code() == codes.InvalidArgument {
		// logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		WriteJSON(w, Response{Error: true,
			Data: Error{
				Status:  http.StatusBadRequest,
				Message: st.Message(),
			}})
		return err

	} else if st.Code() == codes.DataLoss {
		// logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		WriteJSON(w, Response{Error: true,
			Data: Error{
				Status:  http.StatusBadRequest,
				Message: st.Message(),
			}})
		return err

	} else if st.Code() == codes.PermissionDenied {
		// logger.Error(err)
		w.WriteHeader(http.StatusForbidden)
		WriteJSON(w, Response{Error: true,
			Data: Error{
				Status:  http.StatusForbidden,
				Message: st.Message(),
			}})
		return err

	} else if strings.Contains("User blocked in user service", st.Message()) {
		// logger.Error(err)
		w.WriteHeader(http.StatusForbidden)
		WriteJSON(w, Response{Error: true,
			Data: Error{
				Status:  http.StatusForbidden,
				Message: st.Message(),
			}})
		return err
	}
	// logger.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	WriteJSON(w, Response{Error: true,
		Data: Error{
			Status:  http.StatusInternalServerError,
			Message: st.Message(),
		}})

	return err
}

type Response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

// Error ...
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func HandleInternalWithMessage(w http.ResponseWriter, err error, message string) error {
	if err == nil {
		return nil
	}

	log.Panicln(message+" ", err)
	w.WriteHeader(http.StatusInternalServerError)
	WriteJSON(w, Response{Error: true,
		Data: Error{
			Status:  http.StatusInternalServerError,
			Message: message,
		}})
	return err
}

func HandleBadRequestErrWithMessage(w http.ResponseWriter, err error, message string) error {
	if err == nil {
		return nil
	}

	log.Println(message+" ", err)
	w.WriteHeader(http.StatusBadRequest)
	WriteJSON(w, Response{Error: true,
		Data: Error{
			Status:  http.StatusBadRequest,
			Message: message + err.Error(),
		}})
	return err
}

func HandleBadRequestResponse(w http.ResponseWriter, message string) {
	log.Println(message)
	w.WriteHeader(http.StatusBadRequest)
	WriteJSON(w, Response{Error: true,
		Data: Error{
			Status:  http.StatusBadRequest,
			Message: message,
		}})
}

func parsePageQueryParam(r *http.Request) (int, error) {
	pageparam := r.URL.Query().Get("page")
	if pageparam == "" {
		return 1, nil
	}

	page, err := strconv.Atoi(pageparam)
	if err != nil {
		return 0, err
	}
	if page < 0 {
		return 0, errors.New("page must be an positive integer")
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

func parseLimitQueryParam(r *http.Request) (int, error) {

	limitparam := r.URL.Query().Get("limit")
	if limitparam == "" {
		return 10, nil
	}

	page, err := strconv.Atoi(limitparam)
	if err != nil {
		return 0, err
	}
	if page < 0 {
		return 0, errors.New("limit must be an positive integer")
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

func validateUUID(ID string) bool {
	_, err := uuid.Parse(ID)
	return err == nil
}

func BodyParser(r *http.Request, body interface{}) error {
	return json.NewDecoder(r.Body).Decode(&body)
}

func WriteJSON(w http.ResponseWriter, data interface{}) {
	bytes, _ := json.MarshalIndent(data, "", "  ")

	w.Header().Set("Content-Type", "Application/json")
	w.Write(bytes)
}
