package service

import (
	"fmt"
	"time"

	"github.com/SmitaJShetty/rebu/internal/model"
	"github.com/SmitaJShetty/rebu/internal/repo"
)

//CarTripService construct for car service
type CarTripService struct {
	DataRetriever *repo.MedallionRepo
}

//NewCarTripService returns CarTripService
func NewCarTripService() *CarTripService {
	return &CarTripService{
		DataRetriever: repo.NewMedallionRepo(),
	}
}

//GetTripCount handler for fetching trip count
func (c *CarTripService) GetTripCount(medallionNumber string, pickupDate *time.Time) (*model.GetTripResponse, error) {
	if medallionNumber == "" {
		return nil, fmt.Errorf("id was empty in request")
	}

	if pickupDate == nil {
		return nil, fmt.Errorf("invalid pickup date")
	}

	if c.DataRetriever == nil {
		return nil, fmt.Errorf("medallion repo was null")
	}

	trips, getErr := c.DataRetriever.GetTripCount(medallionNumber, pickupDate)
	if getErr != nil {
		return nil, getErr
	}

	return &model.GetTripResponse{
		Response: trips,
	}, nil
}
