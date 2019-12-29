package repo

import (
	"time"

	"github.com/SmitaJShetty/rebu/internal/model"
	"github.com/stretchr/testify/mock"
)

//NewMockMedallionRepo constructor for mock
func NewMockMedallionRepo() *MockMedallionRepo {
	return &MockMedallionRepo{}
}

//MockMedallionRepo construct for repo
type MockMedallionRepo struct {
	mock.Mock
}

//GetTripCount returns trip count
func (m *MockMedallionRepo) GetTripCount(medallionList []string, pickupDate *time.Time) ([]*model.TripSummary, error) {
	return []*model.TripSummary{
		&model.TripSummary{
			Count:      2,
			Medallion:  medallionList[0],
			PickupDate: pickupDate.Format("2006-01-02"),
		},
	}, nil
}
