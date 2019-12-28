package repo

import (
	"time"

	"github.com/stretchr/testify/mock"
)

//MockMedallionRepo construct for repo
type MockMedallionRepo struct {
	mock.Mock
}

//GetTripCount returns trip count
func (m *MockMedallionRepo) GetTripCount(medallionNumber string, pickupDate *time.Time) (int, error) {
	return 20, nil
}
