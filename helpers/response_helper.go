package helpers

import (
	"encoding/json"
	"net/http"
)

type ResponseHelper struct {
}

func (responseHelper ResponseHelper) RespondWithError(w http.ResponseWriter, code int, msg string) {
	responseHelper.RespondWithJSON(w, code, map[string]string{"error": msg})
}

func (responseHelper ResponseHelper) RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
