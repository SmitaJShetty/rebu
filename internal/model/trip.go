package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

//Trip construct for trip
type Trip struct {
	gorm.Model
	ID          string    `json:"id"`
	MedallionID string    `json:"medallion_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

//NewTrip construct for new trip
func NewTrip(medallion string) *Trip {
	return &Trip{
		ID:          uuid.New(),
		MedallionID: medallion,
	}
}
