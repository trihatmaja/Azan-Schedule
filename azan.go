package azan

import (
	"errors"
	"time"
)

// azan
type Azan struct {
	db    DbProvider
	cache CacheProvider
	calc  CalcProvider
}

func New(database DbProvider, cache CacheProvider, calculation CalcProvider) *Azan {
	return &Azan{
		db:    database,
		cache: cache,
		calc:  calculation,
	}
}

func (a *Azan) Generate(latitude, longitude, timezone float64, city string) error {
	v, err := a.db.Validate(latitude, longitude, city)
	if err != nil {
		return err
	}

	if v {
		return errors.New("city with latitude and longitude exists")
	}
	res := a.calc.Calculate(latitude, longitude, timezone, city)

	err = a.db.Set(res)
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

func (a *Azan) GetByCityDate(city string, date time.Time) (CalcResult, error) {
	return a.db.GetByCityDate(city, date)
}

func (a *Azan) GetByCityMonth(city string, month int) (CalcResult, error) {
	return a.db.GetByCityMonth(city, month)
}

func (a *Azan) GetCache(key string) ([]byte, error) {
	return a.cache.Get(key)
}

func (a *Azan) GetCities() ([]CalcResult, error) {
	return a.db.GetCities()
}

func (a *Azan) SetCache(key string, data []byte) error {
	return a.cache.Set(key, data)
}
