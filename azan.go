package azan

import (
	"context"
	"time"
)

type DbInterface interface {
	Set(DbData) (bool, int64)
	GetAll() []DbData
	GetByCity(string) DbData
	GetByDate(string, time.Time) DbData
}

type DbData struct {
	Kota    string
	Bulan   string
	Tanggal time.Time
	Subuh   time.Time
	Fajar   time.Time
	Dhuhur  time.Time
	Ashar   time.Time
	Maghrib time.Time
	Isya    time.Time
}

type CalcInterface interface {
	Calculate()
}

type Azan struct {
	db   DbInterface
	calc CalcInterface
}

func New(database *DbInterface, calculation *CalcInterface) *Azan {
	return &Azan{
		db:   database,
		calc: calculation,
	}
}

func (a *Azan) Generate() {
	a.calc.Calculate()
}
