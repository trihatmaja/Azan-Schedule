package azan

import (
	"time"
)

// database contract
type DbProvider interface {
	Set(CalcResult) error
	GetAll() ([]CalcResult, error)
	GetByCity(string) (CalcResult, error)
	GetByDate(time.Time) (CalcResult, error)
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
	City      string
	Latitude  float64
	Longitude float64
	Timezone  float64
	Schedule  []AzanSchedule
}

type AzanSchedule struct {
	Date    string
	Fajr    string
	Sunrise string
	Zuhr    string
	Asr     string
	Maghrib string
	Isya    string
}
