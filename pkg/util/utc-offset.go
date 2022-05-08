package util

import (
	"fmt"
	"time"

	"github.com/zsefvlol/timezonemapper"
)

func UTCOffset(lat, lon float64) (int, error) { 
	locStr := timezonemapper.LatLngToTimezoneString(lat, lon)

	loc, err := time.LoadLocation(locStr)
	if err != nil {
		return 0, fmt.Errorf("cannot extract utc offset from lat: %f, lon: %f, err: %w", lat, lon, err)
	}
	
	_, offset := time.Now().In(loc).Zone()

	return offset, nil
}
