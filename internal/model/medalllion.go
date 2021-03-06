package model

import (
	"time"

	"github.com/google/uuid"
)

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
	TripDistance    float32   `json:"trip_distance"`
}

//TableName db table name for model Metallion
func (Medallion) TableName() string {
	return "cab_trip_data"
}

//NewMedallion returns a new medallion
func NewMedallion(medNumber string, hlice string, vendorid string, rateCode string, storeFwdFlag string,
	pickupDateTime time.Time, dropOffDateTime time.Time, passCount int, tripTimeSec int64,
	tripDistance float32) *Medallion {
	return &Medallion{
		ID:              uuid.New().String(),
		MedallionNumber: medNumber,
		HackLicense:     hlice,
		VendorID:        vendorid,
		RateCode:        rateCode,
		StoreFwdFlag:    storeFwdFlag,
		PickupDatetime:  pickupDateTime,
		DropoffDatetime: dropOffDateTime,
		PassengerCount:  passCount,
		TripTimeInSecs:  tripTimeSec,
		TripDistance:    tripDistance,
	}
}
