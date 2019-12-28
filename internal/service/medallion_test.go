package service

import (
	"testing"
	"time"

	"github.com/SmitaJShetty/rebu/internal/repo"
	"github.com/stretchr/testify/assert"
	"github.com/oklog/ulid"
)

func Test_GetTripCount(t *testing.T) {
	r := repo.MockMedallionRepo{}
	svc := NewCarTripService(&r)
	t1 := time.Now().AddDate(-1, 0, 0)
	id, idErr := ulid.MustNew(ulid.Timestamp(time.Unix(1000000)),ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0))
	assert.Nil(t, idErr)
	result, err := svc.GetTripCount(id, &t1)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}
