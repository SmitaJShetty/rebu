package service

import (
	"fmt"
	"log"
	"time"

	"github.com/SmitaJShetty/rebu/internal/model"
	"github.com/SmitaJShetty/rebu/internal/repo"
	cachingstore "github.com/SmitaJShetty/rebu/pkg/caching_store"
)

//CarTripService construct for car service
type CarTripService struct {
	DataRetriever repo.DataRetriever
	Cache         cachingstore.CartTripCacheService
}

//NewMockCarTripService returns CarTripService
func NewMockCarTripService() *CarTripService {
	return &CarTripService{
		DataRetriever: repo.NewMockMedallionRepo(),
		Cache:         cachingstore.NewMockCache(),
	}
}

//NewCarTripService returns CarTripService
func NewCarTripService() *CarTripService {
	return &CarTripService{
		DataRetriever: repo.NewMedallionRepo(),
		Cache:         cachingstore.NewCacheService(),
	}
}

//InvalidateCache invalidates cache
func (c *CarTripService) InvalidateCache() error {
	err := c.Cache.ClearCache()
	if err != nil {
		return fmt.Errorf("Error occurred while clearing cache")
	}

	log.Printf("cache was invalidated at %v", time.Now().UTC())
	return nil
}

//GetTripCount handler for fetching trip count
func (c *CarTripService) GetTripCount(medallionList []string, pickupDate *time.Time, isFresh bool) (*model.GetTripResponse, error) {
	if medallionList == nil || len(medallionList) == 0 {
		return nil, fmt.Errorf("id was empty in request")
	}

	if pickupDate == nil {
		return nil, fmt.Errorf("invalid pickup date")
	}

	if c.DataRetriever == nil {
		return nil, fmt.Errorf("medallion repo was null")
	}

	var trips []*model.TripSummary
	var getErr error

	//get from db
	if !isFresh {
		log.Printf("fetching from cache")
		trips, getErr = c.getFromCache(medallionList, pickupDate.Format("2006-01-02"))
		if getErr != nil {
			return nil, getErr
		}

		//All medallions were found in cache
		if trips != nil && len(trips) > 0 {
			log.Println("trips found in cache as requested, exiting with data")
			return &model.GetTripResponse{
				Response: trips,
			}, nil
		}
	}
	//if isFresh is true, or some of the medallions were not found in cache, go to db and refresh cache
	trips, getErr = c.DataRetriever.GetTripCount(medallionList, pickupDate)
	if getErr != nil {
		return nil, getErr
	}

	log.Println("fetched trips from db")
	updateErr := c.updateCache(trips)
	if updateErr != nil {
		return nil, updateErr
	}

	return &model.GetTripResponse{
		Response: trips,
	}, nil
}

//even if one medallion is not present in cache, visit db. Better get fresh data for all if visiting db for one
func (c *CarTripService) getFromCache(medallionList []string, pickupDate string) ([]*model.TripSummary, error) {
	var trips []*model.TripSummary
	for _, s := range medallionList {
		value, getErr := c.Cache.Get(s + "-" + pickupDate)
		if getErr != nil {
			return nil, getErr
		}

		//return with empty trip data
		if value == 0 {
			return nil, nil
		}

		trips = append(trips, &model.TripSummary{
			Medallion:  s,
			PickupDate: pickupDate,
			Count:      value,
		})
	}

	return trips, nil
}

func (c *CarTripService) updateCache(trips []*model.TripSummary) error {
	if trips == nil || len(trips) == 0 {
		log.Println("no trips fetched")
		return nil
	}

	hasError := false
	for _, t := range trips {
		log.Println("updating trip information for trip:", t.Medallion, t.PickupDate)
		updateTripErr := c.updateTripCache(t)
		if updateTripErr != nil {
			hasError = true
			log.Printf("error occurred while updating trip in cache for medallion %s and date %v", t.Medallion, t.PickupDate)
		}
	}

	if hasError {
		return fmt.Errorf(" error occurred while updating trips: %v", trips)
	}
	return nil
}

func (c *CarTripService) updateTripCache(trip *model.TripSummary) error {
	if trip == nil {
		log.Println("trip was empty")
		return nil
	}

	key := trip.Medallion + "-" + trip.PickupDate
	value := trip.Count

	log.Printf("updating key: %s, value: %d", key, value)
	err := c.Cache.Set(key, value)
	return err
}
