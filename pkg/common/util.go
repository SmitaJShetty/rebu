package common

import (
	"fmt"
	"time"
)

//GetDateFromString returns date from string date
func GetDateFromString(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, fmt.Errorf("error: ate string was empty")
	}

	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return nil, fmt.Errorf("error occurred while parsing date %s to format %s, err: %v", dateStr, layout, err)
	}

	return &t, nil
}
