package model

import "../../pkg/apperror"

//GetTripResponse construct for response
type GetTripResponse struct {
	Error    *apperror.AppError `json:"error"`
	Response []*TripSummary     `json:"response"`
}

//NewGetTripResponse returns GetTripResponse
func NewGetTripResponse(message string, statusCode int, response []*TripSummary) *GetTripResponse {
	return &GetTripResponse{
		Error:    apperror.NewAppError(message, statusCode),
		Response: response,
	}
}
