package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetTripCount(t *testing.T) {
	svc := NewCarTripService()
	t1 := time.Now().AddDate(-1, 0, 0)
	id := "01234555"
	result, err := svc.GetTripCount(id, &t1)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}
