package azan

import (
	"errors"
	"time"
)

// database contract
type DbInterface interface {
	Set(CalcResult) error
	GetAll() ([]CalcResult, error)
	GetByCity(string) (CalcResult, error)
	GetByDate(time.Time) (CalcResult, error)
}

// calculation contract
type CalcInterface interface {
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

type ApiRequest struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
	TimeZone  float64 `json:"tz"`
	City      string  `json:"city"`
}

// azan
type Azan struct {
	db   DbInterface
	calc CalcInterface
}

func New(database DbInterface, calculation CalcInterface) *Azan {
	return &Azan{
		db:   database,
		calc: calculation,
	}
}

func (a *Azan) Generate(latitude, longitude, timezone float64, city string) error {
	test, _ := a.db.GetByCity(city)
	if test.City != "" {
		return errors.New("city exist")
	}
	res := a.calc.Calculate(latitude, longitude, timezone, city)

	err := a.db.Set(res)
	if err != nil {
		return err
	}
	return nil
}

func (a *Azan) GetAll() ([]CalcResult, error) {
	return a.db.GetAll()
}

func (a *Azan) GetByCity(city string) (CalcResult, error) {
	return a.db.GetByCity(city)
}
