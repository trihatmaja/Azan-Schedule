package azan

import (
	"log"
	"time"
)

// database contract
type DbInterface interface {
	Set([]CalcResult) error
	GetAll() ([]DbData, error)
	GetByCity(string) (DbData, error)
	GetByDate(time.Time) (DbData, error)
}

type DbData struct {
	City    string `json:"city"`
	Date    string `json:"date"`
	Fajr    string `json:"fajr"`
	Sunrise string `json:"sunrise"`
	Zuhr    string `json:"zuhr"`
	Asr     string `json:"asr"`
	Maghrib string `json:"maghrib"`
	Isya    string `json:"isya"`
}

// calculation contract
type CalcInterface interface {
	Calculate(float64, float64, float64, string) []CalcResult
}

type CalcResult struct {
	City     string
	Month    string
	Year     int
	Schedule []AzanSchedule
}

type AzanSchedule struct {
	Date    int
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
	res := a.calc.Calculate(latitude, longitude, timezone, city)

	err := a.db.Set(res)
	if err != nil {
		return err
	}
	return nil
}

func (a *Azan) GetAll() ([]DbData, error) {
	return a.db.GetAll()
}
