package azan

import (
	"time"
)

// database contract
type DbProvider interface {
	Set(CalcResult) error
	Validate(float64, float64, string) (bool, error)
	GetAll() ([]CalcResult, error)
	GetByCity(string) (CalcResult, error)
	GetByCityDate(string, time.Time) (CalcResult, error)
	GetByCityMonth(string, int) (CalcResult, error)
	GetCities() ([]CalcResult, error)
}

// cache contract
type CacheProvider interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
}

// calculation contract
type CalcProvider interface {
	// latitude, longitude, timezone, city
	Calculate(float64, float64, float64, string) CalcResult
}

type CalcResult struct {
	City      string         `json:"city"`
	Latitude  float64        `json:"lat"`
	Longitude float64        `json:"long"`
	Timezone  float64        `json:"timezone"`
	Schedule  []AzanSchedule `json:"schedule"`
}

type AzanSchedule struct {
	Date    string `json:"date"`
	Fajr    string `json:"fajr"`
	Sunrise string `json:"sunrise"`
	Zuhr    string `json:"zuhr"`
	Asr     string `json:"asr"`
	Maghrib string `json:"maghrib"`
	Isya    string `json:"isya"`
}
