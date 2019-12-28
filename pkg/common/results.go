package common

import (
	"encoding/json"
	"net/http"
)

//SendResult sends result over http response
func SendResult(w http.ResponseWriter, r *http.Request, resultJSON []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}

//SendErrorResponse sends error response
func SendErrorResponse(w http.ResponseWriter, r *http.Request, appErr *AppError) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	json, _ := json.Marshal(appErr)
	w.Write(json)
}
