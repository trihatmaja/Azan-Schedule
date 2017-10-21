package azan

import (
	"log"
	"time"
)

// database contract
type DbInterface interface {
	Set([]CalcResult) error
	GetAll() ([]CalcResult, error)
	GetByCity(string) (CalcResult, error)
	GetByDate(time.Time) (CalcResult, error)
}

// calculation contract
type CalcInterface interface {
	Calculate() []CalcResult
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

func (a *Azan) Generate() {
	res := a.calc.Calculate()

	err := a.db.Set(res)
	if err != nil {
		log.Println(err.Error())
	}

}
