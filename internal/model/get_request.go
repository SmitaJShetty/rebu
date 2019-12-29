package model

import "github.com/SmitaJShetty/rebu/pkg/apperror"

//GetTripResponse construct for response
type GetTripResponse struct {
	Error    *apperror.AppError `json:"error"`
	Response int                `json:"response"`
}

//NewGetTripResponse returns GetTripResponse
func NewGetTripResponse(message string, statusCode int, response int) *GetTripResponse {
	return &GetTripResponse{
		Error:    apperror.NewAppError(message, statusCode),
		Response: response,
	}
}
