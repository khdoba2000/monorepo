package auth_handler

import (
	"encoding/json"
	"net/http"

	"monorepo/src/api_gateway/handlers/models"
)

func writeJSON(w http.ResponseWriter, data interface{}) {
	bytes, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func writeError(w http.ResponseWriter, msg string, code int) {
	bytes, _ := json.Marshal(models.StandartErrorModel{
		ErrorMessage: msg,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bytes)
}
