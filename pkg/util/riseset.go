package util

import (
	"fmt"
	"github.com/arpanetus/kun/pkg/config"
	"github.com/kelvins/sunrisesunset"
	"log"
	"time"
)

// get hour from time
func FormatTime(t time.Time) string {
	return t.Format("15:04:05")
}

type RiseSet struct {
	logger *log.Logger
	config *config.KunConfig
}

func NewRiseSet(logger *log.Logger, config *config.KunConfig) *RiseSet {
	return &RiseSet{
		logger: logger,
		config: config,
	}
}

func (rs *RiseSet) sunriseSunset() (time.Time, time.Time, error) {
	loc, err := Location(rs.config.Latitude, rs.config.Longitude)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("cannot get location: %s", err)
	}
	// let's get the offset from
	ofs, err := UTCOffset(loc)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("cannot get UTC offset: %s", err)
	}

	sunrise, sunset, err := sunrisesunset.GetSunriseSunset(rs.config.Latitude, rs.config.Longitude, float64(ofs), time.Now().In(loc))
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("cannot get sunrise and sunset: %s", err)
	}

	return sunrise, sunset, nil
}

func (rs *RiseSet) SunriseSunset() (string, string, error) {
	rise, set, err := rs.sunriseSunset()
	if err != nil {
		return "", "", fmt.Errorf("cannot get sunrise and sunset: %s", err)
	}

	return FormatTime(rise), FormatTime(set), nil
}

func (rs *RiseSet) IsSunDown() (bool, error) {
	rs.logger.Println("checking if sun is down")

	loc, err := Location(rs.config.Latitude, rs.config.Longitude)
	if err != nil {
		return false, fmt.Errorf("cannot get location: %s", err)
	}

	rise, set, err := rs.sunriseSunset()
	now := time.Now().In(loc)
	if now.After(set) {
		return true, nil
	}

	if now.Before(rise) {
		return true, nil
	}

	return false, nil
}
