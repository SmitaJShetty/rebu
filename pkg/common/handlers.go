package common

import (
	"encoding/json"
	"net/http"

	"github.com/SmitaJShetty/rebu/internal/service"
	"github.com/SmitaJShetty/rebu/pkg/apperror"
	"github.com/gorilla/mux"
)

//Handlers construct for handlers
type Handlers struct {
	CarTripSvc *service.CarTripService
}

//GetTrips handler for GetTrips
func (h *Handlers) GetTrips(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	medallionNumber := vars["medallionNumber"]

	q := req.URL.Query()
	pickupDate := q.Get("pickupdate")

	date, dateErr := GetDateFromString(pickupDate)
	if dateErr != nil {
		SendErrorResponse(w, req, apperror.NewAppError(dateErr.Error(), http.StatusInternalServerError))
		return
	}

	trips, getTripsErr := h.CarTripSvc.GetTripCount(medallionNumber, date)
	if getTripsErr != nil {
		SendErrorResponse(w, req, apperror.NewAppError(getTripsErr.Error(), http.StatusInternalServerError))
		return
	}

	resp, marshalErr := json.Marshal(trips)
	if marshalErr != nil {
		SendErrorResponse(w, req, apperror.NewAppError(marshalErr.Error(), http.StatusInternalServerError))
		return
	}

	if resp == nil {
		SendErrorResponse(w, req, apperror.NewAppError("response was empty", http.StatusInternalServerError))
		return
	}

	SendResult(w, req, []byte(resp))
}

//NewHandlers return new Handlers
func NewHandlers() *Handlers {
	return &Handlers{
		CarTripSvc: service.NewCarTripService(),
	}
}
