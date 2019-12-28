package model

//GetTripResponse construct for response
type GetTripResponse struct {
	Error    *ApiError `json:"error"`
	Response int       `json:"response"`
}

//NewGetTripResponse returns GetTripResponse
func NewGetTripResponse(message string, statusCode int, response string) *GetTripResponse {
	return &GetTripResponse{
		Error:    &common.ApiError{},
		Response: response,
	}
}
