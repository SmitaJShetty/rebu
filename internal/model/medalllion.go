package model

import "time"

//Medallion construct for medallion
type Medallion struct {
	ID              string    `json:"id"`
	MedallionNumber string    `json:"medallion_number"`
	HackLicense     string    `json:"hack_license"`
	VendorID        string    `json:"vendor_id"`
	RateCode        string    `json:"rate_code"`
	StoreFwdFlag    string    `json:"store_fwd_flag"`
	PickupDatetime  time.Time `json:"pickup_datetime"`
	DropoffDatetime time.Time `json:"dropoff_datetime"`
	PassengerCount  int       `json:"passenger_count"`
	TripTimeInSecs  int64     `json:"trip_time_in_secs"`
	TripDistance    int32     `json:"trip_distance"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

//NewMedallion returns a new medallion
func NewMedallion() *Medallion {

}
