package azan

import (
	"context"
	"log"
	"time"
)

type DbInterface interface {
	Set(context.Context, DbData) (bool, int64)
	GetAll(context.Context) []DbData
	GetByCity(context.Context, string) DbData
	GetByDate(context.Context, string, time.Time) DbData
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
