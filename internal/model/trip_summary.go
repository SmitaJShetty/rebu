package model

//TripSummary construct for trip summary
type TripSummary struct {
	Medallion  string `json:"medallion"`
	PickupDate string `json:"pickup_date"`
	Count      int    `json:"count"`
}
