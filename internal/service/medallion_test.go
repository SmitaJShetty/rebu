package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetTripCount(t *testing.T) {
	svc := NewMockCarTripService()
	t1 := time.Now().AddDate(-1, 0, 0)
	medallion := "01234555"
	id := []string{medallion}
	result, err := svc.GetTripCount(id, &t1, false)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, result.Response[0].Count, 2)
	assert.Equal(t, result.Response[0].Medallion, medallion)
}

func Test_GetTripCount_Validation(t *testing.T) {
	svc := NewMockCarTripService()
	t1 := time.Now().AddDate(-1, 0, 0)
	id := []string{}
	result, err := svc.GetTripCount(id, &t1, false)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "id was empty in request")
}

func Test_InvalidateCache(t *testing.T) {
	svc := NewMockCarTripService()
	err := svc.InvalidateCache()
	assert.Nil(t, err)
}
