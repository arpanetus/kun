package util

import (
	"fmt"
	"time"

	"github.com/zsefvlol/timezonemapper"
)

func Location(lat, lon float64) (*time.Location, error) {
	locStr := timezonemapper.LatLngToTimezoneString(lat, lon)
	loc, err := time.LoadLocation(locStr)
	if err != nil {
		return nil, fmt.Errorf("cannot extract utc offset from lat: %f, lon: %f, err: %w", lat, lon, err)
	}
	return loc, nil
}

func UTCOffset(loc *time.Location) (float64, error) {
	_, offset := time.Now().In(loc).Zone()

	offsetInHours := float64(offset / 3600.0)

	return offsetInHours, nil
}
